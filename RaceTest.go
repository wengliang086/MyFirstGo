package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	var lock sync.Mutex

	// 临界资源
	var a = 1
	go func() {
		lock.Lock()
		a = 2
		fmt.Printf("sub go routine %d\n", a)
		lock.Unlock()

		wg.Done()
	}()

	lock.Lock()
	a = 3
	fmt.Println("main go routine ", a)
	lock.Unlock()

	wg.Wait()
}
