package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		p := c.Query("param")
		fmt.Println(p)
		c.String(200, p)
	})

	router.Run(":8080")
}
