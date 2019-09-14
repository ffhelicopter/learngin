package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/user/:firstname/:lastname", func(c *gin.Context) {
		fname := c.Param("firstname")
		lname := c.Param("lastname")

		c.String(http.StatusOK, "Hello %s %s ", fname, lname)
	})

	router.Run(":8080")
}
