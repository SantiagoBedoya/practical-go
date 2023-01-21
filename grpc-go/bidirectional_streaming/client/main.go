package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/SantiagoBedoya/grpc-go/bidirectional_streaming/proto"
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
		fmt.Print("\n> ")
		s.Scan()
		if s.Err() != nil {
			log.Fatal(err)
		}
		stream, err := client.Execute(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			for {
				data, err := stream.Recv()
				if err != nil {
					if err == io.EOF {
						break
					}
					log.Fatal(err)
				}
				log.Printf("Message: %q\n", data.Data)
			}
		}()

		go func() {
			req := &proto.ExecuteRequest{Data: s.Text()}
			if err := stream.Send(req); err != nil {
				log.Fatal(err)
			}
			log.Println("Message sent")
		}()

	}
}
