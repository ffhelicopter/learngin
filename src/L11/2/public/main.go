package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/", "./")
	// 监听并启动服务
	router.Run(":8081")
}
