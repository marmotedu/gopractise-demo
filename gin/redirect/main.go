package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	// 301，永久重定向，跳转到外部链接
	r.GET("/test301", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	// 302，临时重定向，跳转到内部链接
	r.POST("/test302", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/test")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
