package main

import (
	"log"
	"net"

	"github.com/jasmaabox/grpc-tutorial/chat"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port: %s", err)
	}

	grpcServer := grpc.NewServer()

	s := chat.Server{}
	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %s", err)
	}
}
