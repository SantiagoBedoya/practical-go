package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/SantiagoBedoya/grpc-go/bidirectional_streaming/proto"
	"google.golang.org/grpc"
)

type testServer struct{}

func (s *testServer) Execute(srv proto.SomethingService_ExecuteServer) error {
	for {
		req, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		log.Printf("Message received: %q", req.Data)

		res := &proto.ExecuteResponse{Data: fmt.Sprintf("%q processed", req.Data)}
		if err := srv.Send(res); err != nil {
			log.Fatal(err)
		}
	}
	return nil
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
