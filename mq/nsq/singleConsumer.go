package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

type NewHandler struct{}

func (m *NewHandler) HandleMessage(msg *nsq.Message) (err error) {
	addr := msg.NSQDAddress
	message := string(msg.Body)
	fmt.Println(addr, message)
	return
}

func MyConsumers(topic, channel string) {
	conf := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 接收消息
	handler := &NewHandler{}
	consumer.AddHandler(handler)
	err = consumer.ConnectToNSQD(addr)
	if err != nil {
		fmt.Println(err)
	}
}

var addr = "127.0.0.1:4150"

func main() {
	go MyConsumers("topic-demo1", "channel-aa")
	// 模拟多个从多个channel去消息
	go MyConsumers("topic-demo1", "channel-bb")
	select {}
}
