package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/long_async", func(c *gin.Context) {
		// 创建在 goroutine 中使用的副本
		cCp := c.Copy()
		go func() {
			// 用 time.Sleep() 模拟一个长任务。
			time.Sleep(5 * time.Second)

			// goroutine中使用复制的上下文副本 "cCp"
			fmt.Println("完成！ 路径 " + cCp.Request.URL.Path)
		}()
	})

	router.GET("/long_sync", func(c *gin.Context) {
		// 用 time.Sleep() 模拟一个长任务。
		time.Sleep(5 * time.Second)

		// 因为没有不在goroutine中，不需要使用上下文副本
		fmt.Println("完成！ 路径 " + c.Request.URL.Path)
	})

	// 监听并启动服务
	router.Run(":8080")
}
