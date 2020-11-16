package service

import (
	"fmt"
	"golang/client"
	pb "steam_grpc/proto"
)

//StreamService 服务端对象
type StreamService struct {
}

//New 新建服务器对象
func New() *StreamService {
	s := &StreamService{}

	fmt.Printf("New MessageList %p", s)

	return s
}

//Connect 链接
func (s *StreamService) Connect(stream pb.StreamService_ConnectServer) error {
	clientInfo, err := stream.Recv()
	if err != nil {
		return err
	}

	if clientInfo.OperationType == pb.OperationType_LOGIN {
		c := client.New(clientInfo.ServerID, stream)
		return c.Loop()
	}

	fmt.Println("clientInfo.OperationType:", clientInfo.OperationType)
	return nil
}
