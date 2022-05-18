package main

import (
	"context"
	userpb "github.com/luchanos/grpc_app/gen/go/user/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type userService struct {
}

func (u *userService) GetUser(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{
		User: &userpb.User{
			FullName:      "Nikolai Sviridov",
			Uuid:          req.Uuid,
			MartialStatus: userpb.MaritalStatus_MARITAL_STATUS_MARRIED,
			BirthYear:     1991,
			Addresses: []*userpb.Address{
				{
					Street: "Chertanovo",
					City:   "Moscow",
				},
			},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:9879")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &userService{})
	grpcServer.Serve(lis)
}

//type UserServiceServer interface {
//	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
//	mustEmbedUnimplementedUserServiceServer()
//}
