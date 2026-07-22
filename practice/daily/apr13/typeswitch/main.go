package main

import "log"

type MyInterface interface {
	doSomething()
}

type MyStruct struct{}

func (s MyStruct) doSomething() {
	log.Println("hi")
}

func main() {
	var x MyStruct
	// var x int

	typeSwitch(x)

}

func typeSwitch(t any) {
	switch t.(type) {
	case string:
		log.Println("string")
	case int:
		log.Println("int")
	case MyInterface:
		log.Println("My Interface")
	default:
		log.Println("unknown type")
	}

}
