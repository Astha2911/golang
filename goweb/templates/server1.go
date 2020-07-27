package templates
package main


import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates /*.html")
	r.GET("/", displayString)
	r.GET("/user/:name", userInfo)
	r.GET("/html", displayHtml)
	r.POST("/greetings", greet)
	r.Run(":8000")
}
func greet(c *gin.Context) {
	name := c.PostForm("name")
	out := "Hello" + name + "welcome to delhi"
	c.String(http.StatusOK, out)
}
func displayHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}
func userInfo(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"naam hai": name})
}
func displayString(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}
