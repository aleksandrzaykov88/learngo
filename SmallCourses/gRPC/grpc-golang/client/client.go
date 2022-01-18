package main

import (
	"context"
	"fmt"
	"log"

	pb "example.com/grpc/gen/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewTestApiClient(conn)

	resp, err := client.CreateUser(context.Background(), &pb.User{Id: 1, Name: "Alex", Age: 33})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
