package main

import (
	"gingonic-api/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/ping ", handler.PingGet)
	server.Run()

}
