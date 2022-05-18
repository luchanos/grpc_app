package main

import (
	"context"
	"fmt"
	userpb "github.com/luchanos/grpc_app/gen/go/user/v1"
	wearablepb "github.com/luchanos/grpc_app/gen/go/wearable/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
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
	streamingClient := wearablepb.NewWearableServiceClient(conn)
	stream, err := streamingClient.ConsumeBeatsPerMinute(context.Background())

	for i := 0; i < 10; i++ {
		err = stream.Send(&wearablepb.ConsumeBeatsPerMinuteRequest{
			Uuid:   "Nikolai",
			Value:  uint32(i),
			Minute: uint32(i * 2),
		})
		if err != nil {
			log.Fatalln("Send", err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{
		Uuid: "Hello, World!",
	})

	if err != nil {
		log.Fatalf("fail to GetUser: %v", err)
	}

	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln("Close", err)
	}

	fmt.Println(resp.GetTotal())

	fmt.Printf("%+v\n", res)
}
