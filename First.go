package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	var a = 10
	fmt.Printf("value:%d,address:%x\n", a, &a)

	// create Timer
	timeDemo := time.NewTimer(time.Duration(1) * time.Second)
	// 监听定时器
	select {
	case <-timeDemo.C:
		fmt.Println("timer up!")
	}

	fmt.Println("a", "b", 11)
}
