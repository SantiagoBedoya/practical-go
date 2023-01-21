package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/SantiagoBedoya/grpc-go/client_streaming/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewSomethingServiceClient(conn)

	stream, err := client.Execute(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		s.Scan()
		if s.Err() != nil {
			log.Fatal(err)
		}
		str := s.Text()
		if str == "exit" {
			break
		}
		req := &proto.ExecuteRequest{Data: str}
		if err := stream.Send(req); err != nil {
			log.Fatal(err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Resp: %q\n", resp.Data)
}
