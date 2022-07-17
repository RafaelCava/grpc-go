package main

import (
	"context"
	"log"
	"net"

	"github.com/RafaelCava/grpc-go/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.HelloServiceServer
}

func (*server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	res := &pb.HelloResponse{
		Msg: "Salve familiia" + request.GetName(),
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
