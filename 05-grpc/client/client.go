package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.NewClient("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}
	serviceClient := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, err := serviceClient.Add(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.GetResult())
}
