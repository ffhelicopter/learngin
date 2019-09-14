package main

import (
	"net/http"

	"L11/1/pb"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"userinfo": "Mo", "status": http.StatusOK})
	})

	router.GET("/moreJSON", func(c *gin.Context) {
		// 也可使用一个结构体
		var msg struct {
			Name     string `json:"user"`
			UserInfo string
			Number   int
		}
		msg.Name = "Lena"
		msg.UserInfo = "Mo"
		msg.Number = 123
		// 注意 msg.Name 在 JSON 中变成了 "user"
		c.JSON(http.StatusOK, msg)
	})

	router.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"userinfo": "Mo", "status": http.StatusOK})
	})

	router.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"userinfo": "Mo", "status": http.StatusOK})
	})

	router.GET("/someProtoBuf", func(c *gin.Context) {
		name := "Lena"
		// protobuf 的具体定义写在 pb/user文件中。
		data := &pb.UserInfo{
			UserType: 101,
			UserName: name,
			UserInfo: "Mo",
		}
		// 将输出被 protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})

	// 监听并启动服务
	router.Run(":8080")
}
