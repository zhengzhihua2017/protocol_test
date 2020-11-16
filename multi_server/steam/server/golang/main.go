package main

import (
	"golang/module"
	"golang/service"
	"log"
	"net"
	pb "steam_grpc/proto"

	"google.golang.org/grpc"
)

const (
	PORT = "9002"
)

func main() {
	server := grpc.NewServer()
	//服务器启动
	module.Start()

	pb.RegisterStreamServiceServer(server, service.New())

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}
