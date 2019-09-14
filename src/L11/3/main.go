package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 可以自定义添加安全JSON前缀
	// router.SecureJsonPrefix(")]}',\n")

	router.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// SecureJsonPrefix()设置优先，
		// 否则将会输出:   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	// 监听并启动服务
	router.Run(":8080")
}
