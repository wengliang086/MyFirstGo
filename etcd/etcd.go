package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	config := clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:2380"},
		DialTimeout: 5 * time.Second,
	}
	client, err := clientv3.New(config)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer client.Close()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 100*time.Second)
	resp, err := client.Put(ctx, "sample_key", "sample_value")
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("response: %v", resp)
	cancelFunc()
}
