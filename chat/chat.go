package chat

import (
	context "context"
	"log"
)

// Server is gRPC server
type Server struct{}

// SayHello responds to client with a hello message
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message from client: %s", message.Body)
	return &Message{Body: "Server says hi"}, nil
}
