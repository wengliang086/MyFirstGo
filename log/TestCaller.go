package main

import (
	"fmt"
	"runtime"
)

func main() {
	a()
}

func a() {
	b()
}

func b() {
	fmt.Println(runtime.Caller(0))
}
