package main

import (
	"context"
	"fmt"
	"golang/client"
	"io"
	"time"

	pb "steam_grpc/proto"
)

const (
	PORT = "9002"
)

func main() {
	// conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("grpc.Dial err: %v", err)
	// }

	// defer conn.Close()

	// client := pb.NewStreamServiceClient(conn)

	// err = printLists(client, &pb.ClientInfo{
	// 	ServerID:      1,
	// 	OperationType: pb.OperationType_LOGIN,
	// 	Message:       "first login",
	// })
	// if err != nil {
	// 	panic(fmt.Sprintf("printLists.err: %v", err.Error()))
	// }

	c := client.New(":" + PORT)
	c.Login()
	time.Sleep(time.Minute * 60)
}

func printLists(client pb.StreamServiceClient, r *pb.ClientInfo) error {
	stream, err := client.Connect(context.Background())

	stream.Send(r)

	go func() {
		for {
			resp, err1 := stream.Recv()
			if err1 == io.EOF {
				fmt.Println("err == io.EOF")
				break
			}
			if err1 != nil {
				err = err1
				return
			}

			fmt.Println("resp: ", resp.GetMessage())
		}
	}()

	time.Sleep(time.Second * 5)

	go func() {
		for i := 0; i < 100; i++ {
			err2 := stream.Send(&pb.ClientInfo{
				ServerID:      1,
				OperationType: pb.OperationType_CHAT,
				Message:       fmt.Sprintf("chat %d", i),
			})

			if err2 != nil {
				fmt.Println("send : err:", err2.Error())
				return
			}
			time.Sleep(time.Second * 10)
		}
	}()

	time.Sleep(time.Hour)
	if err != nil {
		return err
	}

	return nil
}
