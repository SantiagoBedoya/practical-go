package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/SantiagoBedoya/grpc-go/unary/proto"
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
		resp, err := client.Execute(context.Background(), &proto.ExecuteRequest{Data: s.Text()})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp.Data)
	}
}
