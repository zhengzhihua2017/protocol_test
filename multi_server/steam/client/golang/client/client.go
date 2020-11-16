package client

import (
	"context"
	"fmt"
	"golang/input"
	"golang/view"
	pb "steam_grpc/proto"

	"google.golang.org/grpc"
)

//Client 客户端对象
type Client struct {
	*grpc.ClientConn
}

//New 新建客户端
func New(address string) *Client {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("grpc.Dial err: %v", err))
	}

	c := &Client{}
	c.ClientConn = conn

	return c
}

//Login 登陆
func (c *Client) Login() {
	streamServiceClient := pb.NewStreamServiceClient(c.ClientConn)
	stream, err := streamServiceClient.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	id := input.GetID()
	err = stream.Send(&pb.ClientInfo{
		ServerID:      id,
		OperationType: pb.OperationType_LOGIN,
		Message:       fmt.Sprintf("%d login", id),
	})

	if err != nil {
		panic(err)
	}

	go view.Run(stream)
	go input.Run(stream)
}

//Close 关闭
func (c *Client) Close() {
	c.ClientConn.Close()
}
