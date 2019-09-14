package main

import (
	"fmt"
	"net/http"

	_ "github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func BindHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// c.ShouldBind 使用了 c.Request.Body，不可重用。
	if errA := c.ShouldBind(&objA); errA != nil {
		fmt.Println(errA)
		c.String(http.StatusOK, `the body should be formA`)
		// 因为现在 c.Request.Body 是 EOF，所以这里会报错。
	} else if errB := c.ShouldBind(&objB); errB != nil {

		fmt.Println(errB)
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		c.String(http.StatusOK, `Success`)
	}
}

/*
func BindHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// ShouldBindBodyWith() 读取 c.Request.Body 并将结果存入上下文。
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA != nil {

		fmt.Println(errA)
		c.String(http.StatusOK, `the body should be formA`)
		// 这时, 复用存储在上下文中的 body。
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB != nil {
		fmt.Println(errB)
		c.String(http.StatusOK, `the body should be formB JSON`)
		// 可以接受其他格式
	} else {
		c.String(http.StatusOK, `Success`)
	}
}
*/
func main() {
	route := gin.Default()
	route.Any("/bind", BindHandler)
	route.Run(":8080")
}
