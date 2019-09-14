package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	// router.LoadHTMLFiles("templates/index.tmpl", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	// 监听并启动服务
	router.Run(":8080")
}
