package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ryanralphs/blog/data"

	"github.com/gin-gonic/gin"
)

func main() {
	fakeData := data.FakePosts(10)
	allPosts, _ := json.Marshal(fakeData)
	svr := gin.Default()
	svr.SetTrustedProxies([]string{"localhost"})
	svr.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})
	svr.GET("/posts/", func(c *gin.Context) {
		c.String(http.StatusOK, string(allPosts))
	})

	svr.GET("/posts/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		c.String(http.StatusOK, string(fakeData[id].Title)+" : "+string(fakeData[id].Content))
	})
	svr.Run("localhost:8080")
}
