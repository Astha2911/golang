package main

import (
	"asthaapi/controllers"
	"asthaapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Initializing a gin router
	//r.LoadHTMLGlob("html/*")

	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	//})
	models.ConnectDatabase()
	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.LoginUser)
	r.Run()
}
