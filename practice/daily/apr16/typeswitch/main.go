package main

import "log"

type MyInterface interface {
	doSomething()
}

type MyStruct struct {
}

func (s *MyStruct) doSomething() {}

func typeSwitch(t any) {
	switch t.(type) {
	case MyInterface:
		log.Println("my interface")
	case string:
		log.Println("string")
	case int:
		log.Println("int")
	default:
		log.Println("unknown type")
	}
}

func main() {
	var s *MyStruct

	typeSwitch(s)
}
