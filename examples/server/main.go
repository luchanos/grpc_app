package main

import (
	"context"
	userpb "github.com/luchanos/grpc_app/gen/go/user/v1"
)

type userService struct {
}

func (u *userService) GetUser(context.Context, *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return nil, nil
}

func main() {

}

//type UserServiceServer interface {
//	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
//	mustEmbedUnimplementedUserServiceServer()
//}
