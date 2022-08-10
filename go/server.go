package main

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
    "fmt"
	"os"
	
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context){c.String(200, "pong")})
	router.GET("/search", searchAPI)
	router.GET("/notionOAuth", notionOAuthAPI)
	return router
}

func searchAPI(c *gin.Context) {
	searchUseNotionAPI()
	c.String(200, "search")
}

type RequestNotion struct {
	Query string `json:"query"`
	Sort SortOption `json:"sort"`
}

type SortOption struct {
	Direction string `json:"direction"`
	Timestamp string `json:"timestamp"`
}

type GetDataParser struct {
	Object string `json:"object"`
	id string `json:"id"`
}

func searchUseNotionAPI() {
	baseURL := "https://api.notion.com/v1"
	jsonString, err := json.Marshal(RequestNotion{Query: "ソフトウェア", Sort: SortOption{Direction: "descending", Timestamp: "last_edited_time"}})
	if err != nil {
        panic("Error")
    }
	
	req, err := http.NewRequest("POST", baseURL+"/search", bytes.NewBuffer(jsonString))
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
	fmt.Printf("%#v", string(byteArray))

	// get page data
	var data []GetDataParser
	err = json.Unmarshal(byteArray, &data)
	if err != nil {
        fmt.Println(err)
    }
	fmt.Println(data)
	/*

	req, err = http.NewRequest("GET", baseURL+"/pages/"+)
	if err != nil {
        panic("Error")
    }

	f, _ = os.Open("secure")
	token = make([]byte, 64)
	n,_ = f.Read(token)
	req.Header.Set("Authorization", "Bearer "+string(token[:n]))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Content-Type", "application/json")

	client = new(http.Client)
	resp, err = client.Do(req)
	if err != nil {
		fmt.Print(err)
        panic("Error")
    }

	defer resp.Body.Close()

	byteArray, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        panic("Error")
    }

    fmt.Printf("%#v", string(byteArray))*/
}

func notionOAuthAPI(c *gin.Context) {
	c.String(200, "Oauth")
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}