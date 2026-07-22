package main

import "log"

type Interface interface {
	doSomething()
}

type SomeStruct struct {
}

func (s *SomeStruct) doSomething() {
	log.Print("doing somet")
}

func main() {
	// s := &SomeStruct{}
	// s := 1
	s := "how are you"
	typeSwitch(s)
}

func typeSwitch(v any) {
	switch v.(type) {
	case Interface:
		log.Println("type found")
	case int:
		log.Println("int found")
	case string:
		log.Println("string found")
	default:
		log.Println("type not found")
	}
}
