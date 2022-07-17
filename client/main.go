package main

import (
	"context"
	"log"

	"github.com/RafaelCava/grpc-go/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	Hello(err, client)
}

func Hello(err error, client pb.HelloServiceClient) {
	request := &pb.HelloRequest{
		Name: "Rafael",
	}

	res, err := client.Hello(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Println(res.Msg)
}
