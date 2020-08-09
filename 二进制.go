package main

import (
	"encoding/json"
	"fmt"
)

var intArr [14]int16

func init() {
	for i := 0; i < len(intArr); i++ {
		intArr[i] = 1 << i
	}
	fmt.Println(intArr)
}

func main() {
	marshal, _ := json.Marshal([]int{1, 1, 1})
	fmt.Printf("数组转Json: %s\n", string(marshal))

	str := "[1,1,1,1,0,0,0,0,0]"
	var arr []int
	err := json.Unmarshal([]byte(str), &arr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Json转数组: %v\n", arr)

	var total int
	for i := 0; i < len(arr); i++ {
		a := 1 << (len(arr) - 1 - i)
		//fmt.Println(a)
		if arr[i] == 1 {
			total += a
		}
	}
	fmt.Printf("数组转Int: %d\n", total)

	fmt.Printf("Int转二进制：")
	v := int16(total)
	for i := len(arr) - 1; i >= 0; i-- {
		if v >= intArr[i] {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
		v -= intArr[i]
	}

	var bitmap = []int16{480, 480, 480, 480, 0, 0, 0, 0, 0}
	markedMap := bitMapTo(bitmap, 100, 100)
	fmt.Println(markedMap)
}

func bitMapTo(bitmap []int16, tx, ty int32) map[int32]map[int32]bool {
	marked := make(map[int32]map[int32]bool)

	markCoordinate := func(marked map[int32]map[int32]bool, x int32, y int32) bool {
		row := marked[x]
		exists := false
		if row == nil {
			row = make(map[int32]bool)
		}

		exists = row[y]
		if !exists {
			row[y] = true
			marked[x] = row
		}
		return !exists
	}

	for i, v := range bitmap {
		x := tx - int32(i)
		y := ty + int32(i)
		for i := len(bitmap) - 1; i >= 0; i-- {
			if v >= intArr[i] {
				markCoordinate(marked, x, y)
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}

			v -= intArr[i]
			x += 1
			y += 1
		}
		fmt.Println()
	}

	return marked
}
