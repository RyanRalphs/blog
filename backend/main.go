package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ryanralphs/blog/database"
	"github.com/ryanralphs/blog/structs"

	"github.com/gin-gonic/gin"
)

func main() {

	collection, ctx, err := database.ConnectDB("blog-posts")

	if err != nil {
		log.Fatal(err)
	}
	
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
	svr.POST("/posts", func(c *gin.Context) {
		var post structs.Post
		c.BindJSON(&post)
		fakeData = append(fakeData, post)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		fmt.Println(fakeData)

	})
	svr.Run("localhost:8080")
}
