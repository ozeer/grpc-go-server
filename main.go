package main

import (
	"context"
	"fmt"
	"grpc-go-server/service"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	service.UnimplementedUserServer
}

func (s *server) Register(ctx context.Context, in *service.RegisterReq) (*service.RegisterResp, error) {
	return &service.RegisterResp{Code: 200, Message: "成功", Data: ""}, nil
}

func (s *server) Login(ctx context.Context, in *service.LoginReq) (*service.LoginResp, error) {
	return &service.LoginResp{AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXkiOiJhcHAiLCJleHAiOjE2ODU1OTMzMDQsIm5iZiI6MTY4NTU0OTEwNCwiaWF0IjoxNjg1NTUwMTA0LCJqdGkiOiI0NjMwODUyOTg3MjE0ODg5MDEifQ.7S-IQu-F4JpMZo1s9WFQYGA9j5EqfuoSir-JmQKj5QU", ExpireIn: 12600}, nil
}

func main() {
	// 监听本地端口
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		fmt.Printf("监听端口失败: %s", err)
		return
	}
	// 创建gRPC服务器
	s := grpc.NewServer()
	// 注册服务
	service.RegisterUserServer(s, &server{})
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}
}
