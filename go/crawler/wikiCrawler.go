package main

import (
	"fmt"
	"net/http"
	"sort"
	"regexp"
	"strings"
	"github.com/PuerkitoBio/goquery"

	"database/sql"
	"github.com/jinzhu/copier"
	"github.com/bluele/mecab-golang"
	_"github.com/mattn/go-sqlite3"
)

func getIntListDiff(listA []int, listB []int) ([]int) {
	// listAに存在してlistBに存在しない要素の取得関数
	largeList := listA
	smallList := listB

	// ソートして調べる
	sort.Slice(largeList,func(i, j int) bool { return largeList[i] < largeList[j] })
	sort.Slice(smallList,func(i, j int) bool { return smallList[i] < smallList[j] })
	isSameList := make([]bool, len(largeList)) // init all false
	for i,_ := range largeList {
		for j := 0; j < len(smallList); j++ {
			if smallList[j] == largeList[i] {
				isSameList[i] = true
				break
			}
		}
	}
	var diffList []int
	for i,isSame := range isSameList {
		if !isSame {
			diffList = append(diffList, largeList[i])
		}
	}
	return diffList
}

func isContain(list []int, target int) bool {
	for _,v := range list {
		if v==target {
			return true
		}
	}
	return false
}

// <---- modeling ---->
const (
	db_path = "./"
	db_name = "test.db"
)

func DBinit() (error){
	// 正常に生成出来たらnilを返す
	db, err := sql.Open("sqlite3", db_path+db_name)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		`CREATE TABLE "words" ("id" integer primary key not null, "word" text)`,
	)
	_, err = db.Exec(
		`CREATE TABLE "article_words" ("word_id" integer not null, "article_id" integer not null, primary key(article_id,word_id))`,
	)
	if err != nil {
		if _,err = db.Exec(`delete from words`);err!=nil {
			return err
		}
		if _,err = db.Exec(`delete from article_words`);err!=nil {
			return err
		}
	}
	return nil	
}

func getAllWordInArticle() (string, error) {
	db, err := sql.Open("sqlite3", db_path+db_name)
	if err != nil {
		return "", err
	}
	defer db.Close()

	res, err := db.Query(
		`select * from article_words`,
	)
	if err != nil {
		return "", err
	}
	for res.Next() {
		var word_id int
		var article_id int
		if err := res.Scan(&word_id, &article_id);err!=nil {
			return "",err
		}
		fmt.Printf("word_id:%d, article_id:%d\n", word_id,article_id)
	}
	return "", err
}

func getWordID(word string) (id int,err error) {
	db, err := sql.Open("sqlite3", db_path+db_name)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	res := db.QueryRow(
		`SELECT * FROM words WHERE word=?`,
		word,
	)
	err = res.Scan(&id)
	if err!=nil {
		return nil, err
	}
	return
}

func getArticleIDs(wordID int) (articleIDs []int, err error) {
	db, err := sql.Open("sqlite3", db_path+db_name)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	res, err := db.Query(
		`select article_id from article_words where word_id=?`,
		wordID,
	)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		var id int
		if err := res.Scan(&id);err!=nil {
			return nil, err
		}
		articleIDs = append(articleIDs, id)
	}
	return
}

func insertNewWord(word string, articleIDs []int) (error){
	db, err := sql.Open("sqlite3", db_path+db_name)
	if err != nil {
		return err
	}
	defer db.Close()

	// 単語を追加する
	res, err := db.Exec(
		`insert into words (word) values (?)`,
		word,
	)
	if err != nil {
		return err
	}
	// 追加したidを取得して記事と結びつける
	wordID,err := res.LastInsertId()
	if err != nil {
		return err
	}
	_, err = db.Exec(
		`insert into article_words (word_id, article_id) values (?,?)`,
		wordID,
		articleIDs[0],
	)
	if err!= nil {
		return err
	}
	return nil
}

func updateArticleID(word string,articleIDs []int) (error){
	db, err := sql.Open("sqlite3", db_path+db_name)
	if err != nil {
		return err
	}
	defer db.Close()

	res := db.QueryRow(
		`SELECT id FROM words WHERE word=?`,
		word,
	)

	var wordID int
	err = res.Scan(&wordID)
	if err!=nil {
		return err
	}
	
	var articleIDsLOG []int
	articleIDsLOG, err = getArticleIDs(wordID)
	if err!=nil {
		return err
	}

	diff := getIntListDiff(articleIDs, articleIDsLOG)
	for _,articleId := range diff {
		_, err = db.Exec(
			`insert into article_words (word_id,article_id) values (?,?)`,
			wordID,
			articleId,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

//<---- crawler ---->

func get_a_article(url string)string {
	resp, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	/*
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	*/

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil{
		panic(err)
	}

	retrunText := ""

	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		rawText := s.Text()
		
		re := regexp.MustCompile("[[0-9]+]")
		context := re.ReplaceAllString(rawText, "")
		re = regexp.MustCompile("[0-9]+年")
		context = re.ReplaceAllString(context, "")
		re = regexp.MustCompile("\n")
		context = re.ReplaceAllString(context, "")
		retrunText += context
	})
	return retrunText
}

func scrape() string {
	urlList := [2]string{
		"https://ja.wikipedia.org/wiki/%E6%9C%88%E3%83%8E%E7%BE%8E%E5%85%8E", //月ノ美兎
		"https://ja.wikipedia.org/wiki/%E5%90%8D%E5%8F%96%E3%81%95%E3%81%AA", //名取さな
	}

	for i := range urlList {
		return get_a_article(urlList[i])
	}
	return ""
}

//<---- word controller ---->
func splitArticleToSentence(article string) []string {
	return strings.Split(article,"。")
}

func splitSentenceToWords(sentence string) ([]string, error) {
	m,err := mecab.New("-Owakati")
	if err != nil {
		return nil,err
	}
	defer m.Destroy()

	tag, err := m.NewTagger()
	if err != nil {
		return nil,err
	}
	defer tag.Destroy()

	lt, err := m.NewLattice(sentence)
	if err != nil {
		return nil,err
	}
	defer lt.Destroy()

	var wordList []string
	node := tag.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(),",")
		if features[0] != "BOS/EOS" {
			wordList = append(wordList,node.Surface())
		}
		if node.Next() != nil {
			break
		}
	}
	return wordList, nil
}

//<---- sentence controller ---->
func getSentence()

//<---- search ---->
type articleScore struct {
	articleID int
	score int
}

type wordScore struct {
	wordID int
	score int
}

func getScore(inputText string) (score []int) {
	// 入力テキストとDB内の
	score = fullTextSearch(inputText)
	return
}

func wordTranslateToID(word string) (id int) {
	data := dbData{word: word}
	id = data.id
	return
}

func fullTextSearch(text string) (score []int) {
	//words, err := splitSentenceToWords(text)
	_, err := splitSentenceToWords(text)
	if err != nil {
		panic(err)
	}
	/*
	var wordIDs []int
	var wordStructs []wordScore
	maxIndex := 0
	for i,word := range words {
		var data dbData
		data.id = i
		data.word = word
		wordStructs = append(wordStructs, wordScore{wordID: data.id, score: score})
	}*/
	score = [] int{1, 1}
	return
}

func main() {
	if err := DBinit(); err!=nil {
		panic(err)
	}
}