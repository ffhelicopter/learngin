package main

import (
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func main() {
	router := gin.Default()
	router.GET("/user", func(c *gin.Context) {
		// 生成令牌并将其传递到CSRF标头中。
		// JSON客户端或JavaScript框架读取头信息得到令牌。
		// 在随后的请求中自定应请求头“X-CSRF-Token”写入令牌信息。
		c.Header("X-CSRF-Token", csrf.Token(c.Request))
		c.JSON(http.StatusOK, csrf.Token(c.Request))
	})

	// 监听并启动服务

	s := &http.Server{
		Addr:           ":8080",
		Handler:        csrf.Protect([]byte("32-byte-long-auth-key"))(router),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
