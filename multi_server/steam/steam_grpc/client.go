// package simple_grpc

// import (
// 	"context"
// 	"fmt"
// 	"simple_grpc/proto"
// 	"time"

// 	"google.golang.org/grpc"
// )

// //SimpleClient grpc客户端
// type SimpleClient struct {
// 	proto.MessageTransferClient
// 	*grpc.ClientConn
// }

// //NewSimpleClient 新建客户端
// func NewSimpleClient(address string) *SimpleClient {
// 	c := &SimpleClient{}
// 	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
// 	if err != nil {
// 		panic(fmt.Sprintf("did not connect: %v", err))
// 	}
// 	c.ClientConn = conn

// 	c.MessageTransferClient = proto.NewMessageTransferClient(c.ClientConn)

// 	return c
// }

// //Send 发送消息
// func (c *SimpleClient) Send(message string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	r, err := c.MessageTransferClient.Send(ctx, &proto.Message{Message: message})
// 	if err != nil {
// 		panic(fmt.Sprintf("failed to serve: %v", err))
// 	}

// 	fmt.Println(r.GetMessage())
// }

// //Close 关闭
// func (c *SimpleClient) Close() {
// 	c.ClientConn.Close()
// }
