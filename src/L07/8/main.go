package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// Recovery 中间件接收panic异常,并返回500状态码。
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		panic("内部错误!!!")
		c.String(200, "pong")
	})

	router.GET("/test", func(c *gin.Context) {
		c.String(200, "test")
	})
	router.Run(":8080")
}
