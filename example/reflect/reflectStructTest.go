package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	School string
}

func main() {
	s1 := Student{"zhanSan", 19, "qianfeng"}
	fmt.Printf("%T\n", s1)
	p1 := &s1
	fmt.Printf("%T\n", p1)
	fmt.Println(s1.Name)
	fmt.Println((*p1).Name, p1.Name) // 指针直接调用是语法糖

	fmt.Println("--------------")
	// 改变
	value1 := reflect.ValueOf(s1)
	fmt.Println(value1.Kind(), value1.Type())

	value := reflect.ValueOf(p1)
	fmt.Println(value.Kind(), value.Type())
	if value.Kind() == reflect.Ptr {
		newValue := value.Elem()
		fmt.Println(newValue.CanSet())

		f1 := newValue.FieldByName("Name")
		f1.SetString("hanru")
		f3 := newValue.FieldByName("School")
		f3.SetString("幼儿园")
		fmt.Println(s1)
	}
}
