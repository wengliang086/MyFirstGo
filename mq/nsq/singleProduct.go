package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)

func main() {
	addr := "127.0.0.1:4150"
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(addr, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		msg := "message :" + time.Now().Format("2006-01-02 15:04:05")
		err := producer.Publish("topic-demo1", []byte(msg))
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(time.Second * 2)
	}
}
