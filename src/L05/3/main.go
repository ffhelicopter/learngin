package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NoResponse 请求的url不存在，返回404
func NoResponse(c *gin.Context) {
	// 返回404状态码
	c.String(http.StatusNotFound, "404, page not exists!")
}

func main() {
	// 正式发布模式
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// 设定请求url不存在的返回值
	router.NoRoute(NoResponse)

	router.Run(":8080")
}
