package main

import (
	"context"
	test "goModTest/tests/1grpc/pb"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	requestStr := "test111"
	if len(os.Args) > 1 { // os.Args[1] 为用户执行输入的参数 如：go run ***.go 123
		requestStr = os.Args[1]
	}
	client := test.NewWaiterClient(conn)
	res, err := client.DoMD5(context.Background(), &test.Req{
		JsonStr:              requestStr,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("服务器响应：%s", res.BackJson)
}
