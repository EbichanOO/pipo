package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	router.GET("/search", searchAPI)
	router.GET("/notionOAuth", notionOAuthAPI)
	return router
}

func searchAPI(c *gin.Context) {
	searchUseNotionAPI("ソフトウェア")
	c.String(200, "search")
}

type RequestNotion struct {
	Query string     `json:"query"`
	Sort  SortOption `json:"sort"`
}

type SortOption struct {
	Direction string `json:"direction"`
	Timestamp string `json:"timestamp"`
}

type ReturnNotionData struct {
	Object string          `json:"object"`
	Result []GetDataParser `json:"results"`
}

type GetDataParser struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

/* return json 2022/08/13
{
	\"object\":\"list\",
	\"results\": [{
		\"object\":\"page\",
		\"id\":\"cb37d04f-92be-4311-8ed7-7262fba56338\",
		\"created_time\":\"2022-06-28T14:05:00.000Z\",
		\"last_edited_time\":\"2022-06-28T14:05:00.000Z\",
		\"created_by\":{
			\"object\":\"user\",
			\"id\":\"8697a8e9-7f9c-45b7-9cb8-4a962d69da58\"
		},
		\"last_edited_by\":{
			\"object\":\"user\",
			\"id\":\"8697a8e9-7f9c-45b7-9cb8-4a962d69da58\"},
			\"cover\":null,
			\"icon\":null,
			\"parent\":{
				\"type\":\"database_id\",
				\"database_id\":\"31eb4cd6-320f-4e92-a326-8af7927b5853\"
			},
			\"archived\":false,
			\"properties\":{
				\"タイムスタンプ\":{\"id\":\"%3Bj%3Fs\"},
				\"URL\":{\"id\":\"%3ClpL\"},
				\"タグ\":{\"id\":\"fIoT\"},
				\"タイトル\":{\"id\":\"title\"}
			},
		\"url\":\"https://www.notion.so/cb37d04f92be43118ed77262fba56338\"
	}],
	\"next_cursor\":null,
	\"has_more\":false,
	\"type\":\"page_or_database\",
	\"page_or_database\":{}
}
*/

func searchUseNotionAPI(targetWord string) {
	// notionで検索APIを利用して内容を取得する
	// 2022-08-13現在はタイトルマッチのみ実装されている
	
	// get search
	baseURL := "https://api.notion.com/v1"
	jsonString, err := json.Marshal(RequestNotion{
		Query: targetWord, 
		Sort: SortOption{Direction: "descending", Timestamp: "last_edited_time"},
	})
	if err != nil {
		panic("Error")
	}

	req, err := http.NewRequest("POST", baseURL+"/search", bytes.NewBuffer(jsonString))
	if err != nil {
		panic("Error")
	}

	f, _ := os.Open("secure")
	token := make([]byte, 64)
	n, _ := f.Read(token)
	req.Header.Set("Authorization", "Bearer "+string(token[:n]))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
		panic("Error")
	}

	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Error")
	}
	// fmt.Printf("%#v\n\n", string(byteArray))

	// get a page data
	var data ReturnNotionData
	err = json.Unmarshal(byteArray, &data)
	if err != nil {
		fmt.Println(err)
	}
	
	if(len(data.Result) == 0) {
		fmt.Println(" no data matched ...")
	}

	for _, result := range data.Result {
		if(result.Object == "page") {
			fmt.Println("this is page mode")
			fmt.Println(getNotionPageData(result.Id))
		}
	}
}

type ParseBlockData struct {
	Child struct {
		Title string `json:"title"`
	} `json:"child_page"`
}

func getNotionPageData(id string) (string){
	baseURL := "https://api.notion.com/v1"
	req, err := http.NewRequest("GET", baseURL+"/pages/"+id, nil)
	if err != nil {
		panic("Error")
	}

	f, _ := os.Open("secure")
	token := make([]byte, 64)
	n,_ := f.Read(token)
	req.Header.Set("Authorization", "Bearer "+string(token[:n]))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
		panic("Error")
	}

	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Error")
	}

	var data ParseBlockData
	err = json.Unmarshal(byteArray, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(byteArray))

	return data.Child.Title
}

func notionOAuthAPI(c *gin.Context) {
	c.String(200, "Oauth")
}

func main() {
	searchUseNotionAPI("ガジェット")
	// test
	router := setupRouter()
	router.Run(":8080")
}
