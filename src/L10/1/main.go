package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	route.POST("/testing", startPage)
	route.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
	// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，
	// 然后再使用 `Form`（`form-data`）。
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
	}

	c.String(200, "Success")
}
