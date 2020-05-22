package main

import (
	"context"
	"log"

	"github.com/jasmaabox/grpc-tutorial/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Client says hi!!!",
	}

	res, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error sending SayHello: %s", err)
	}

	log.Printf("Response from server: %s", res.Body)
}
