package main

import "testing"

// https://blog.csdn.net/weixin_34232617/article/details/91854391
// go test -bench=. -run=none
// go test -bench=. -run=none -benchtime=3s
// go test -bench=. -benchmem -cpuprofile profile.out
// go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out

type channel chan int

func NoDefer() {
	ch := make(channel, 10)
	close(ch)
}

func Defer() {
	ch := make(channel, 10)
	defer close(ch)
}

func BenchmarkNoDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoDefer()
	}
}

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Defer()
	}
}
