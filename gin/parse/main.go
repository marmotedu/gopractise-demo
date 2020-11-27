package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   string `uri:"id" form:"id" json:"id" header:"id"`
	Name string `uri:"name" form:"name" json:"name" header:"name"`
}

func main() {
	r := gin.Default()

	// 示例 1：获取路径参数
	// 客户端：
	// curl http://127.0.0.1:8080/users/colin/12
	// 输出：
	// {"Param":"colin:12","ShouldBindUri":"colin:12","fullpath":"/users/:name/:id","type":"path"}
	r.Any("/users/:name/:id", func(c *gin.Context) {
		// 获取指定的路径参数
		name := c.Param("name")
		id := c.Param("id")

		// 绑定路径参数，使用 tag：`uri`
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"type":          "path",
			"Param":         fmt.Sprintf("%s:%s", name, id),
			"ShouldBindUri": fmt.Sprintf("%s:%s", person.Name, person.ID),
			"fullpath":      c.FullPath(),
		})
	})

	// 示例 2：获取查询字符串参数
	// 客户端：
	// curl 'http://127.0.0.1:8080/query?name=colin&id=12'
	// 输出：
	// {"Query":"colin:12","ShouldBindQuery":"colin:12","type":"query"}
	r.Any("/query", func(c *gin.Context) {
		// 获取指定的查询字符串参数
		name := c.DefaultQuery("name", "jeremy")
		id := c.Query("id")

		// 绑定查询字符串参数，使用 tag：`form`
		var person Person
		if err := c.ShouldBindQuery(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"type":            "query",
			"Query":           fmt.Sprintf("%s:%s", name, id),
			"ShouldBindQuery": fmt.Sprintf("%s:%s", person.Name, person.ID),
		})
	})

	// 示例 3：表单参数
	// 客户端：
	// curl -F "name=colin" -F "id=12" http://127.0.0.1:8080/form
	// 输出：
	// {"PostForm":"colin:12","type":"form"}
	r.Any("/form", func(c *gin.Context) {
		// 获取指定的表单参数
		name := c.DefaultPostForm("name", "jeremy")
		id := c.PostForm("id")

		c.JSON(200, gin.H{
			"type":     "form",
			"PostForm": fmt.Sprintf("%s:%s", name, id),
		})
	})

	// 示例 4：消息体参数
	// 客户端：
	// curl -XPOST -d'{"name":"colin", "id":"12"}' http://127.0.0.1:8080/body
	// 输出：
	// {"ShouldBindJSON":"colin:12","type":"body"}
	r.Any("/body", func(c *gin.Context) {
		// 绑定消息体参数，使用 tag：`json`
		var person Person
		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"type":           "body",
			"ShouldBindJSON": fmt.Sprintf("%s:%s", person.Name, person.ID),
		})
	})

	// 示例 5：获取 HTTP 头参数
	// 客户端：
	// curl -H "name: colin" -H "id: 12" http://127.0.0.1:8080/header
	// 输出：
	// {"GetHeader":"colin:12","ShouldBindHeader":"colin:12","type":"header"}
	r.GET("/header", func(c *gin.Context) {
		// 获取指定的 HTTP 头参数
		name := c.GetHeader("name")
		id := c.GetHeader("id")

		// 绑定 HTTP 头参数，使用 tag：`header`
		var person Person
		if err := c.ShouldBindHeader(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"type":             "header",
			"GetHeader":        fmt.Sprintf("%s:%s", name, id),
			"ShouldBindHeader": fmt.Sprintf("%s:%s", person.Name, person.ID),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
