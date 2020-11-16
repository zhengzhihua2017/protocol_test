package client

import (
	"fmt"
	"golang/module"
	"io"
	pb "steam_grpc/proto"
)

//Client 客户端
type Client struct {
	id    int32
	steam pb.StreamService_ConnectServer
}

//New 新建客户端对象
func New(id int32, steam pb.StreamService_ConnectServer) *Client {
	c := &Client{
		id:    id,
		steam: steam,
	}

	module.Register(c)

	c.steam.SendMsg(&pb.Message{Message: "login success"})
	return c
}

//Loop 循环
func (c *Client) Loop() error {
	defer func() {
		fmt.Println("loop end")
	}()

	ctx := c.steam.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			err := c.receive()
			if err != nil && err != io.EOF {
				return err
			}
		}
	}
}

func (c *Client) receive() error {
	clientInfo, err := c.steam.Recv()
	if err != nil {
		fmt.Println("receive err", err.Error())
	}

	if clientInfo.GetMessage() != "" {
		module.BroadcastMessage(clientInfo.GetMessage())
	}

	return err
}

//Send 发送消息
func (c *Client) Send(msg string) error {
	fmt.Println("send msg:", msg)
	err := c.steam.Send(&pb.Message{Message: msg})
	return err
}

//GetID 获取玩家ID
func (c *Client) GetID() int32 {
	return c.id
}
