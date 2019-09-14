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
		cc, err := c.Cookie("UserCookie")

		if err != nil || cc == "" {
			errcookie := map[string]interface{}{
				"foo": "Not Cookie",
			}
			c.SetCookie("UserCookie", "Gin", 0, "/", "www.a.com", false, false)
			c.JSONP(http.StatusOK, errcookie)
		} else {
			// callback 是 x
			// 将输出：x({\"foo\":\"bar\"})
			c.JSONP(http.StatusOK, data)
		}
	})

	// 监听并启动服务
	router.Run(":8080")
}
