package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/SantiagoBedoya/grpc-go/unary/proto"
	"google.golang.org/grpc"
)

type testServer struct{}

func (s *testServer) Execute(ctx context.Context, req *proto.ExecuteRequest) (*proto.ExecuteResponse, error) {
	return &proto.ExecuteResponse{
		Data: fmt.Sprintf("%q processed", req.Data),
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", "localhost:5001")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterSomethingServiceServer(grpcServer, &testServer{})

	log.Println("gRPC server is listening on port 5001")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
