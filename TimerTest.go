package main

import (
	"fmt"
	"time"
)

func testTimer1() {
	go func() {
		fmt.Println("test timer1")
	}()
}

func testTimer2() {
	go func() { fmt.Println(time.Now().String()) }()
}

func timer1() {
	time1 := time.NewTimer(1 * time.Second)
	for {
		select {
		case <-time1.C:
			testTimer1()
		}
	}
}

func timer2() {
	time2 := time.NewTimer(5 * time.Second)
	for {
		select {
		case <-time2.C:
			testTimer2()
		}
	}
}

func main() {
	go timer1()
	//timer2()
}
