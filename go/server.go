package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context){c.String(200, "pong")})
	router.GET("/search", searchAPI)
	router.GET("/notionOAuth", notionOAuthAPI)
	return router
}

func searchAPI(c *gin.Context) {
	c.String(200, "search")
}

func notionOAuthAPI(c *gin.Context) {
	c.String(200, "Oauth")
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}