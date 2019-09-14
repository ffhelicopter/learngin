package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("GinCookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("GinCookie", "test", 3600, "/", "localhost", false, true)
		}
		s := fmt.Sprintf("Cookie value: %s \n", cookie)
		c.String(200, s)
	})

	router.Run(":8080")
}
