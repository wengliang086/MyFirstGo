package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

type NewHandler struct {
	channel string
}

func (m *NewHandler) HandleMessage(msg *nsq.Message) (err error) {
	addr := msg.NSQDAddress
	message := string(msg.Body)
	fmt.Println(addr, m.channel, message)
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
	handler := &NewHandler{channel: channel}
	consumer.AddHandler(handler)
	err = consumer.ConnectToNSQD(addr)
	if err != nil {
		fmt.Println(err)
	}
}

var addr = "127.0.0.1:4150"

func main() {
	// 模拟多个从多个channel去消息
	go MyConsumers("topic-demo1", "channel-dd")
	select {}
}
