package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err == nil {
			fmt.Printf("Cookie value: %s\n", cookie)
		}

		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	})

	router.Run()
}
