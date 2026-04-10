package main

import (
	"fmt"
)

type MyInterface interface {
	myfunc() string
}
type MyStruct struct{}

func (ms MyStruct) myfunc() string {
	return ""
}

func main() {
	// var t MyInterface
	var t MyStruct

	typeSwitch(t)
}

func typeSwitch(t any) {
	switch t.(type) {
	case MyInterface:
		fmt.Println("my interface")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("dont know")
	}
}
