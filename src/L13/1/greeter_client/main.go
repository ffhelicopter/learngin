package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	pb "gin/chapter-6/6.2/4/helloworld"
	"google.golang.org/grpc"
)

func main() {
	// 连接gRPC服务端
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	// gin框架HTTP服务设置
	r := gin.Default()
	r.GET("/rest/n/:name", func(c *gin.Context) {
		name := c.Param("name")

		// 调用gRPC服务的方法并把返回通过HTTP响应显示
		req := &pb.HelloRequest{Name: name}
		res, err := client.SayHello(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res.Message),
		})
	})

	// 启动HTTP服务
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
