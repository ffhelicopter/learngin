package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{

			// 做为附件下载保存
			// "Content-Disposition": `attachment; filename="gopher.png"`,

			// 直接在浏览器显示
			"Content-Type": "image/png",
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	router.GET("/file", func(c *gin.Context) {
		// 输出当前目录main.go文件内容
		c.File("main.go")
	})

	router.GET("/attach", func(c *gin.Context) {
		// 下载当前目录main.go文件为test-main.go
		c.FileAttachment("./main.go", "test-main.go")
	})

	router.Run(":8080")
}
