package main

import "log"

type I interface {
	doSomething()
}

type S struct{}

func (s *S) doSomething() {
	log.Println("something")
}

func typeSwitch(t any) {
	switch t.(type) {
	case I:
		log.Println("I")
	case int:
		log.Println("int")
	case string:
		log.Println("string")
	default:
		log.Println("unknown")
	}
}

func main() {
	// var v *S
	var v string
	typeSwitch(v)
}
