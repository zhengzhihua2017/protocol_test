package main

import (
	"context"
	"log"
	"os"
	"server/proto"
	"time"

	"google.golang.org/grpc"
)

const (
	address        = "localhost:50051"
	defaultMessage = "hello world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewChatClient(conn)

	// Contact the server and print out its response.
	msg := defaultMessage
	if len(os.Args) > 1 {
		msg = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Send(ctx, &proto.ChatMessage{Message: msg})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("client: %s", r.GetMessage())
}
