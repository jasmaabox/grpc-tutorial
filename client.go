package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

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
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("<: ")

		s, _ := reader.ReadString('\n')

		res, err := c.SayHello(context.Background(), &chat.Message{
			Body: s,
		})
		if err != nil {
			log.Fatalf("Error sending SayHello: %s", err)
		}

		fmt.Printf(">: %s\n", res.Body)
	}
}
