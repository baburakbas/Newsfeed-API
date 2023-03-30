package main

import (
	"gingonic-api/httpd/handler"
	newsfeed "gingonic-api/platform/newsFeed"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()

	server := gin.Default()

	server.GET("/ping", handler.PingGet())
	server.GET("/newsfeed", handler.NewsFeedGet(feed))
	server.POST("/newsfeed", handler.NewsFeedPost(feed))

	server.Run()

}
