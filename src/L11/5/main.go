package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Unicode实体
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 纯字符标识
	router.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 监听并启动服务
	router.Run(":8080")
}
