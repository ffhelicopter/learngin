package main

import (
	"github.com/gin-gonic/gin"
)

type CheckForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	router := gin.Default()

	router.Static("/", "./public")

	router.POST("/check", func(c *gin.Context) {
		var form CheckForm

		// 简单地使用 ShouldBind 方法自动绑定
		if c.ShouldBind(&form) == nil {
			c.JSON(200, gin.H{"color": form.Colors})
		}
	})
	router.Run(":8080")
}
