package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// https://my.oschina.net/solate/blog/3034188
// go tool pprof -http=":8002" http://localhost:8001/debug/pprof/profile

func main() {
	http.HandleFunc("/", sayHelloWorld) // 设置访问路由
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	helloWorld(10000)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Fprintf(w, "Hello world!\n")
}

func helloWorld(times int) {
	time.Sleep(time.Second)
	var counter int
	for i := 0; i < times; i++ {
		for j := 0; j < times; j++ {
			counter++
		}
	}
}
