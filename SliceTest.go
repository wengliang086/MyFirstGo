package main

import "fmt"

func main() {
	var myArray []int
	myArray = make([]int, 10, 1000)

	myArray[0] = 1
	myArray[1] = 2
	//myArray[100] = 100
	fmt.Println(len(myArray))

	for _, v := range myArray {
		fmt.Println(v)
	}

	var myArray2 [10]int
	myArray2[0] = 1
	myArray2[9] = 9
	//myArray2[10] = 10
	fmt.Println(len(myArray2))
	for i := 0; i < len(myArray2); i++ {
		fmt.Println(myArray2[i])
	}
}
