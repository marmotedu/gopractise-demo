package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个不带任何中间件的路由
	r := gin.New()

	// 全局中间件
	// Logger 中间件将日志写到 gin.DefaultWriter，即使设置了 GIN_MODE=release
	// 默认设置 gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件，从任何 panic 恢复，并返回一个 500 错误
	r.Use(gin.Recovery())

	// 对于每一个路由，如果有需要，可以添加多个中间件
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 授权组
	// authorized := r.Group("/", AuthRequired())
	// 也可以这样
	authorized := r.Group("/")
	// 在这个示例中，我们使用了一个自定义的中间件 AuthRequired()，该中间件只作用于 authorized 组
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// 嵌套组
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// 监听并服务于0.0.0.0:8080
	r.Run(":8080")
}
