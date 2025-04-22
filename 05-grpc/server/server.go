package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type AppServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

func (asi *AppServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("Processing x = %d, y = %d\n", x, y)
	time.Sleep(5 * time.Second)
	result := x + y
	fmt.Printf("Sending result = %d\n", result)
	return &proto.AddResponse{
		Result: result,
	}, nil
}

func (asi *AppServiceImpl) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	fmt.Printf("Generating Prime number from %d to %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			fmt.Printf("Sending Prime No : %d\n", no)
			if err := serverStream.Send(res); err != nil {
				log.Fatalln(err)
			}
			time.Sleep(300 * time.Millisecond)
		}
	}
	return nil
}

func (asi *AppServiceImpl) Aggregate(serverStream proto.AppService_AggregateServer) error {
	var sum, min, max int64 = 0, 9223372036854775807, -9223372036854775808
LOOP:
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			log.Println("[AppService - Aggregate] All the data have been received")
			res := &proto.AggregateResponse{
				Sum: sum,
				Min: min,
				Max: max,
			}
			if err := serverStream.SendAndClose(res); err != io.EOF && err != nil {
				log.Fatalln(err)
			}
			break LOOP
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(req)
		no := req.GetNo()
		sum += no
		if no < min {
			min = no
		}
		if no > max {
			max = no
		}
	}
	return nil
}

func isPrime(no int64) bool {
	for i := int64(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	asi := &AppServiceImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)

	grpcServer.Serve(listener)
}
