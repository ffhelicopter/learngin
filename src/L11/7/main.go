package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 直接重定向到外部站点
	router.GET("/re", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	// 间接重定向到内部Handler
	router.GET("/handle", func(c *gin.Context) {
		c.Request.URL.Path = "/hi"
		router.HandleContext(c)
		c.JSON(200, gin.H{"Handler": "handle"})
	})

	router.GET("/hi", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	// 监听并启动服务
	router.Run(":8080")
}
