package main

import (
	"fmt"
	"reflect"
)

/**
 * 指定长度或者 中括号 中是 ... 的就是数组，否则就是切片
 */
func main() {
	var myArray []int
	myArray = make([]int, 10, 1000)

	myArray[0] = 1
	myArray[1] = 2
	//myArray[100] = 100
	fmt.Println(len(myArray))
	fmt.Printf("len=%d cap=%d slice=%v\n", len(myArray), cap(myArray), myArray)
	myArray = append(myArray, 11, 22, 33)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(myArray), cap(myArray), myArray)

	for _, v := range myArray {
		fmt.Println(v)
	}

	var myArray2 [10]int
	myArray2[0] = 1
	myArray2[9] = 9
	//myArray2[10] = 10
	fmt.Println(len(myArray2))
	fmt.Printf("len=%d cap=%d slice=%v\n", len(myArray2), cap(myArray2), myArray2)
	//myArray2 = append(myArray2, 111)
	for i := 0; i < len(myArray2); i++ {
		fmt.Println(myArray2[i])
	}

	fmt.Println(reflect.TypeOf(myArray), reflect.TypeOf(myArray2))
}
