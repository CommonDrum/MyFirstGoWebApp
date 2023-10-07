package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.POST("/repeat", func(c *gin.Context) {
		message := c.PostForm("message")
		c.HTML(200, "repeat.html", gin.H{
			"message": message,
		})
	})

	r.GET("/next", func(c *gin.Context) {
		c.HTML(200, "next_page.html", nil)
	})

	r.Run() // listens on port 8080 by default
}
