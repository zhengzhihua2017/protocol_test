// package simple_grpc

// import (
// 	"fmt"
// 	"net"
// 	"simple_grpc/proto"

// 	"google.golang.org/grpc"
// )

// //SimpleServer grpc服务器
// type SimpleServer struct {
// 	*grpc.Server
// 	net.Listener
// }

// //NewSimpleServer 新建服务
// func NewSimpleServer(address string, server proto.MessageTransferServer) *SimpleServer {
// 	s := &SimpleServer{}

// 	lis, err := net.Listen("tcp", address)
// 	if err != nil {
// 		panic(fmt.Sprintf("failed to listen: %v", err))
// 	}
// 	s.Listener = lis

// 	s.Server = grpc.NewServer()
// 	proto.RegisterMessageTransferServer(s.Server, server)
// 	return s
// }

// //Start 服务启动
// func (s *SimpleServer) Start() {
// 	if err := s.Server.Serve(s.Listener); err != nil {
// 		panic(fmt.Sprintf("failed to serve: %v", err))
// 	}
// }

// //Stop 服务关闭
// func (s *SimpleServer) Stop() {
// 	s.Server.Stop()
// }
