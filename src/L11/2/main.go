package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// callback 是 x
		// 将输出：x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	// 监听并启动服务
	router.Run(":8080")
}
