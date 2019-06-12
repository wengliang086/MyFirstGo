package main

import "fmt"

type Phone interface {
	call()
	name()
}

type Nokia struct {
}

func (nokia Nokia) call() {
	fmt.Printf("I am Nokia, I can call you!\n")
}

func (nokia Nokia) name() {
	fmt.Printf("My name is Nokia\n")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Printf("I am iPhone, I can call you!\n")
}

func (iPhone IPhone) name() {
	fmt.Printf("My name is IPhone\n")
}

func main() {
	var phone Phone

	phone = new(Nokia)
	phone.call()
	phone.name()

	phone = new(IPhone)
	phone.call()
	phone.name()
}
