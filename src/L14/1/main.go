package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://google.com"}

	//router.Use(cors.New(config))

	router.Use(cors.Default())

	router.GET("/getuser", func(c *gin.Context) {
		c.String(http.StatusOK, "get")
	})

	router.PUT("/putuser", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "put",
		}
		c.JSON(http.StatusOK, data)
	})

	router.Run(":8080")
}
