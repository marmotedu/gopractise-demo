package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func readFile() (string, error) {
	f, err := ioutil.ReadFile("/tmp/delve.log")
	if err != nil {
		return "", err
	}

	return string(f), nil
}

func main() {
	r := gin.Default()
	r.GET("/content", func(c *gin.Context) {
		log.Println("ping func called.")
		content, err := readFile()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unknown error"})
			return
		}

		c.String(200, string(content))
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
