package input

import (
	"bufio"
	"os"
	pb "steam_grpc/proto"
	"strconv"
)

//Scanner 输入扫描
type Scanner struct {
	*bufio.Scanner
}

var gScanner *bufio.Scanner

func init() {
	gScanner = bufio.NewScanner(os.Stdin)
}

//Run 运行
func Run(stream pb.StreamService_ConnectClient) {
	for {
		gScanner.Scan()
		msg := gScanner.Text()
		stream.Send(&pb.ClientInfo{
			OperationType: pb.OperationType_CHAT,
			Message:       msg,
		})
	}
}

//GetID 获取id
func GetID() int32 {
	gScanner.Scan()
	msg := gScanner.Text()
	id, err := strconv.ParseInt(msg, 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(id)
}
