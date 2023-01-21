package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/SantiagoBedoya/grpc-go/server_streaming/proto"
	"google.golang.org/grpc"
)

type testServer struct{}

func (s *testServer) Execute(req *proto.ExecuteRequest, srv proto.SomethingService_ExecuteServer) error {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(count int) {
			defer wg.Done()
			time.Sleep(time.Duration(count) * time.Second)
			resp := &proto.ExecuteResponse{
				Data: fmt.Sprintf("%q processed", req.Data),
			}
			if err := srv.Send(resp); err != nil {
				log.Fatal(err)
			}
			log.Println("Message sent")
		}(i)
	}

	wg.Wait()
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
