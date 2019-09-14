package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		fmt.Println(person.Name)
		fmt.Println(person.Address)
		c.String(200, "Success")
	} else {
		c.String(400, "Error")
	}

}

func main() {
	route := gin.Default()
	route.Any("/bindquery", startPage)
	route.Run(":8080")
}
