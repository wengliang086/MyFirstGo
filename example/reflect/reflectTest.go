package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
	Desc string
}

func (s *MyStruct) GetName() string {
	return s.name
}

func main() {
	fmt.Println("--------------")
	var a MyStruct
	b := new(MyStruct)

	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.ValueOf(b))
	fmt.Println(reflect.ValueOf(*b))

	fmt.Println("--------------")
	a.name = "aaa"
	b.name = "bbb"
	fmt.Println(reflect.ValueOf(a))
	fmt.Println(reflect.ValueOf(b))
	fmt.Println("--------------")

	fieldA := reflect.ValueOf(a).FieldByName("name")
	fmt.Printf("引用结构体字段值 %v\n", fieldA)
	fieldB := reflect.ValueOf(*b).FieldByName("name")
	fmt.Printf("指针结构体字段值 %v\n", fieldB)

	oldType := reflect.ValueOf(b) // 获取指针的 reflect.Type
	newType := oldType.Elem()     // 获取type的真实类型
	newType.FieldByName("Desc").SetString("rDesc")
	nameField := newType.FieldByName("name")
	if nameField.CanSet() {
		nameField.SetString("rName")
	}

	fmt.Println(b)

	fmt.Println("--------------")

	//value := reflect.ValueOf(b).Elem()
	//value

	fmt.Println("--------------")
}
