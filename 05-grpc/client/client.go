package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"time"

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
	// doRequestResponse(ctx, serviceClient)
	// doServerStreaming(ctx, serviceClient)
	doClientStreaming(ctx, serviceClient)

}

func doRequestResponse(ctx context.Context, serviceClient proto.AppServiceClient) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, err := serviceClient.Add(timeoutCtx, req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.GetResult())
}

func doServerStreaming(ctx context.Context, serviceClient proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		Start: 2,
		End:   100,
	}
	clientStream, err := serviceClient.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All the generated prime numbers have been received")
			break
		}
		if err == nil {
			fmt.Printf("Prime No : %d\n", res.GetPrimeNo())
			continue
		}
		log.Fatalln(err)
		break
	}
}

func doClientStreaming(ctx context.Context, serviceClient proto.AppServiceClient) {
	nos := []int64{3, 1, 4, 2, 5, 9, 6, 8, 7}
	clientStream, err := serviceClient.Aggregate(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		fmt.Println("Sending no :", no)
		req := &proto.AggregateRequest{
			No: no,
		}
		if err := clientStream.Send(req); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("Client finished sending all the data")
	if res, err := clientStream.CloseAndRecv(); err == io.EOF || err == nil {
		fmt.Println("Sum :", res.GetSum())
		fmt.Println("Min :", res.GetMin())
		fmt.Println("Max :", res.GetMax())
	} else {
		log.Fatalln(err)
	}
}
