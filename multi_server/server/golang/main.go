package main

import (
	"context"
	"fmt"
	"log"
	"simple_grpc"
	"simple_grpc/proto"
	"time"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	proto.UnimplementedMessageTransferServer
}

// Send implements helloworld.GreeterServer
func (s *server) Send(ctx context.Context, msg *proto.Message) (*proto.Message, error) {
	log.Printf("Received: %v", msg.GetMessage())
	return &proto.Message{Message: "server " + msg.GetMessage()}, nil
}

func main() {
	fmt.Println("step1")
	simpleServer := simple_grpc.NewSimpleServer(port, &server{})
	fmt.Println("step2")
	simpleServer.Start()
	fmt.Println("simple server start")
	time.Sleep(time.Minute)
}
