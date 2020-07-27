package handler

import (
	"goapi/plateform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsfeedGet(feed *newsfeed.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "found me",
		})
	}
}
