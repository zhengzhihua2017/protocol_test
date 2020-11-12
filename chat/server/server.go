package main

import (
	"context"
	"log"
	"net"
	"server/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	proto.UnimplementedChatServer
}

// Send implements helloworld.GreeterServer
func (s *server) Send(ctx context.Context, msg *proto.ChatMessage) (*proto.ChatReply, error) {
	log.Printf("Received: %v", msg.GetMessage())
	return &proto.ChatReply{Message: "server " + msg.GetMessage()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterChatServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
