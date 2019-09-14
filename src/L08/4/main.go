package main

import (
	"time"

	"github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(limit.MaxAllowed(1))

	router.GET("/test", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(200, "test")
	})

	router.Run(":8080")
}
