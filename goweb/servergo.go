package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", displayString)
	r.Run(":1313")
}
func displayString(c *gin.Context) {
	c.String(http.StatusOK, "hello world") // what we have to display
}
