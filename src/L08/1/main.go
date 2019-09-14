package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("LoggerMiddle Strat: ", time.Now())

		// 转到下一个Handler
		c.Next()

		// Handler已经处理完，继续本中间件的处理
		// 显示Handler响应的状态码以及内容长度
		fmt.Println(c.Writer.Status(), ":", c.Writer.Size())
		fmt.Println("LoggerMiddle End: ", time.Since(t))

	}
}

func main() {
	router := gin.New()

	// 加入自定义中间件
	router.Use(LoggerMiddle())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// 监听并启动服务
	router.Run(":8080")
}
