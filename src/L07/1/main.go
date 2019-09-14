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

	// 加入全局中间件日志中间件
	router.Use(gin.Logger())

	router.GET("/login", Login)

	router.Run(":8080")
}
