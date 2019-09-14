package main

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.String(200, "Login")
}

func Sub(c *gin.Context) {
	c.String(200, "Sub")
}

func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	// 简单的组路由: v1
	v1 := router.Group("/v1", gin.Logger())
	{
		v1.GET("/login", Login)
		v1.GET("/submit", Sub)
	}

	// 简单的组路由: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/log", Login)
		v2.GET("/sub", Sub)
	}

	router.Run(":8080")
}
