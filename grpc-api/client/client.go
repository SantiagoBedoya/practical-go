package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/SantiagoBedoya/grpc-api/gen/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client := pb.NewTestServiceClient(conn)

	resp, err := client.Echo(context.Background(), &pb.EchoRequest{Msg: "Hello world"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
