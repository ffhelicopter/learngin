package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 模拟数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
	router := gin.Default()

	// 组路由使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// "localhost:8080/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，由 BasicAuth 中间件设置
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// 监听并启动服务
	router.Run(":8080")
}
