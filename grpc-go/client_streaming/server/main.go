package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/SantiagoBedoya/grpc-go/client_streaming/proto"
	"google.golang.org/grpc"
)

type testServer struct{}

func (s *testServer) Execute(srv proto.SomethingService_ExecuteServer) error {
	result := ""
	for {
		data, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				return srv.SendAndClose(&proto.ExecuteResponse{Data: result})
			}
			return err
		}
		result += fmt.Sprintf("%q\n", data.Data)
		log.Println(data.Data)
	}
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
