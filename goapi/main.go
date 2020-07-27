package main

import (
	"goapi/httpd/handler"
	"goapi/plateform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet())
	r.POST("/newsfeed", handler.NewsfeedPost())
	r.Run()

}
