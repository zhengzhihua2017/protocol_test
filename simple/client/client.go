package main

import (
	"context"
	"fmt"
	"log"
	"protocol_test/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":6666", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error:%v\n", err)
	}

	defer conn.Close()

	//实例化 UserInfoService 微服务的客户端
	client := proto.NewUserInfoServiceClient(conn)

	//调用服务
	req := new(proto.UserRequest)
	req.Name = "liang"
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("resp error:%v\n", err)
	}

	fmt.Printf("Received: %v\n", resp)

}
