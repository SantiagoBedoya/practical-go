package main

import (
	"context"
	"log"
	"net"

	pb "github.com/SantiagoBedoya/grpc-api/gen/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type testServer struct {
}

func (s *testServer) Echo(context.Context, *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Uuid: uuid.NewString()}, nil
}
func (s *testServer) GetUser(context.Context, *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, nil
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTestServiceServer(grpcServer, &testServer{})

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
