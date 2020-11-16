package main

import (
	"simple_grpc"
)

const (
	address        = "localhost:50051"
	defaultMessage = "hello world"
)

func main() {
	sc := simple_grpc.NewSimpleClient(address)
	sc.Send(defaultMessage)
}
