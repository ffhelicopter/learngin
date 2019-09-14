package main

import (
	"context"
	"log"
	"net"

	pb "gin/chapter-6/6.2/4/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server 用来实现 helloworld.GreeterServer
type server struct{}

// SayHello方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	// 在gRPC服务上反射注册服务
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
