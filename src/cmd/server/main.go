package main

import (
	"context"
	"fmt"
	pb "koizumi55555/grcp/src/pkg/grpc"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserServer struct {
	pb.UnimplementedUsersServiceServer
}

func NewUsersServer() *UserServer {
	return &UserServer{}
}

func main() {
	// 9000番portのLisnterを作成
	port := 9000
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	// gRPCサーバーを作成
	s := grpc.NewServer()
	reflection.Register(s)

	// gRPCサーバーにGreetingServiceを登録
	pb.RegisterUsersServiceServer(s, NewUsersServer())

	// 作成したgRPCサーバーを、9000番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}

func (c *UserServer) GetUser(
	ctx context.Context, req *pb.UserRequest,
) (*pb.UserResponse, error) {
	log.Println("gRpc server's GetUser was called")
	res := pb.UserResponse{
		Id:       "12345",
		UserName: "test taro",
		Email:    "user1@test.com",
	}

	return &res, nil
}

func (c *UserServer) GetUsers(
	ctx context.Context, req *pb.UsersRequest,
) (*pb.UsersResponse, error) {
	log.Println("gRpc server's GetUsers was called")
	users := []*pb.UserResponse{
		{
			Id:       "12345",
			UserName: "test taro",
			Email:    "user1@test.com",
		},
		{
			Id:       "23456",
			UserName: "test jiro",
			Email:    "user2@test.com",
		},
	}

	return &pb.UsersResponse{
		UserList: users,
	}, nil
}
