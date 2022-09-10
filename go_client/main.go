package main

import (
	"context"
	"go_client/greet"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// don't forget to update the address
	conn, err := grpc.Dial("localhost:80", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	greeterClient := greet.NewGreeterClient(conn)
	replyMessage, err := greeterClient.SayHello(context.TODO(), &greet.HelloRequest{
		Name: "Leo",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", replyMessage.GetMessage())
}
