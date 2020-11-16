package module

import (
	"fmt"
	"sync"
)

type manager struct {
	memberList  sync.Map
	messageList chan string
}

type client interface {
	Send(msg string) error
	GetID() int32
}

var gManager *manager

func init() {
	gManager = &manager{
		messageList: make(chan string, 1000),
	}
}

//Register 注册客户端
func Register(c client) {
	Unregister(c.GetID())
	gManager.memberList.Store(c.GetID(), c)
}

//Unregister 注销客户端
func Unregister(id int32) {
	gManager.memberList.Delete(id)
}

//Start 开启服务
func Start() {
	go func() {
		for {
			select {
			case msg := <-gManager.messageList:
				broadcastMessage(msg)
			}
		}
	}()
}

func broadcastMessage(msg string) {
	gManager.memberList.Range(func(key, value interface{}) bool {
		c := value.(client)
		err := c.Send(msg)
		if err != nil {
			fmt.Printf("broadcastMessage err:%s  id:%d", err.Error(), c.GetID())
			Unregister(c.GetID())
		}
		return true
	})
}

//BroadcastMessage 广播信息
func BroadcastMessage(msg string) {
	gManager.messageList <- msg
}
