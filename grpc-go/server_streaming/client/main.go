package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/SantiagoBedoya/grpc-go/server_streaming/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewSomethingServiceClient(conn)

	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		s.Scan()
		if s.Err() != nil {
			log.Fatal(err)
		}
		req := &proto.ExecuteRequest{Data: s.Text()}
		stream, err := client.Execute(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		done := make(chan bool)
		go func() {
			for {
				res, err := stream.Recv()
				if err != nil {
					if err == io.EOF {
						done <- true
						return
					}
					log.Fatal(err)
				}
				log.Printf("Resp received: %q\n", res.Data)
			}
		}()
		<-done
	}
}
