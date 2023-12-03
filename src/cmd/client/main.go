package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "koizumi55555/grcp/src/pkg/grpc"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client pb.UsersServiceClient
)

func main() {
	fmt.Println("start gRPC Client.")

	// gRPCサーバーとのコネクションを確立
	address := "localhost:9000"
	conn, err := grpc.Dial(
		address,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}

	// gRPCクライアントを生成
	client = pb.NewUsersServiceClient(conn)

	ctx := context.Background()
	router := gin.Default()
	router.GET("/v1/users/:id", func(c *gin.Context) {
		getUser(ctx, c, client)
	})
	router.GET("/v1/users", func(c *gin.Context) {
		getUsers(ctx, c, client)
	})
	router.Run(":8080")
}

func getUser(ctx context.Context, c *gin.Context, client pb.UsersServiceClient) {
	log.Println("gRpc client's GetUser was called")
	id := c.Param("id")
	corpPb := pb.UserRequest{
		Id: id,
	}

	res, err := client.GetUser(context.Background(), &corpPb)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, &res)
	err = json.NewEncoder(os.Stdout).Encode(&res)
	if err != nil {
		fmt.Println(err)
	}
}

func getUsers(ctx context.Context, c *gin.Context, client pb.UsersServiceClient) {
	log.Println("gRpc client's GetUsers was called")
	corpsPb := pb.UsersRequest{}
	res, err := client.GetUsers(context.Background(), &corpsPb)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, &res)
	err = json.NewEncoder(os.Stdout).Encode(&res)
	if err != nil {
		fmt.Println(err)
	}
}
