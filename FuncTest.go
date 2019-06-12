package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func multiReturn(a int) (int, int) {
	return a * a, a * a * a
}

func main() {
	println(max(1, 2))
	fmt.Println(max(4, 3))

	i, i2 := multiReturn(3)
	fmt.Println(i, i2)
}
