package main

import (
	"context"
	"fmt"
	userpb "github.com/luchanos/grpc_app/gen/go/user/v1"
	wearablepb "github.com/luchanos/grpc_app/gen/go/wearable/v1"
	"google.golang.org/grpc"
	"io"
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
	wearableService := &wearableService{}
	userpb.RegisterUserServiceServer(grpcServer, &userService{})
	wearablepb.RegisterWearableServiceServer(grpcServer, wearableService)
	grpcServer.Serve(lis)
}

func (w *wearableService) ConsumeBeatsPerMinute(stream wearablepb.WearableService_ConsumeBeatsPerMinuteServer) error {
	var total int
	for {
		value, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wearablepb.ConsumeBeatsPerMinuteResponse{
				Total: uint32(total),
			})
		}
		if err != nil {
			return err
		}
		fmt.Println(value.GetMinute(), value.GetUuid(), value.GetValue())
		total++
	}
}
