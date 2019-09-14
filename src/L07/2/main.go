package main

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.String(200, "Login")
}

func main() {
	// 新建一个没有任何默认中间件的路由
	router := gin.New()

	// 加入全局中间件错误捕获中间件
	router.Use(gin.Recovery())

	// 简单的组路由 v1，直接加入日志中间件
	v1 := router.Group("/v1", gin.Logger())
	{
		v1.GET("/login", Login)
	}

	// 简单的组路由 v2
	v2 := router.Group("/v2")
	// 使用Use()方法加入日志中间件
	v2.Use(gin.Logger())
	{
		v2.GET("/login", Login)
	}

	router.Run(":8080")
}
