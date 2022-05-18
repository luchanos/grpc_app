package main

import (
	"context"
	"fmt"
	userpb "github.com/luchanos/grpc_app/gen/go/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9879", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()
	client := userpb.NewUserServiceClient(conn)

	res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{
		Uuid: "Hello, World!",
	})

	if err != nil {
		log.Fatalf("fail to GetUser: %v", err)
	}

	fmt.Printf("%+v\n", res)
}
