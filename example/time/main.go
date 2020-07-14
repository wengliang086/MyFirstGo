package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Time{})
	fmt.Println(time.Now())
	fmt.Println(time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local))
	fmt.Println(time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05"))
	fmt.Println(time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00"))
}
