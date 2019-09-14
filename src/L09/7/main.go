package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v \n", ids, names)
	})
	router.Run(":8080")
}
