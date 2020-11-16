package view

import (
	"fmt"
	"io"
	pb "steam_grpc/proto"
)

//Run 运行
func Run(stream pb.StreamService_ConnectClient) {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("err == io.EOF")
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Println("recv: ", resp.GetMessage())
	}
}
