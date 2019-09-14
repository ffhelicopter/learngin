package main

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.String(200, "Login")
}

func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	// 为单个处理程序添加任意数量的中间件。
	router.GET("/login", gin.Logger(), Login)

	router.Run(":8080")
}
