package main

import "fmt"

func main() {
	var a = 20
	var ip = &a
	fmt.Printf("a var address is %x\n", &a)
	fmt.Printf("ip stored value address is %x\n", ip)
	fmt.Printf("*ip value is %d\n", *ip)
}
