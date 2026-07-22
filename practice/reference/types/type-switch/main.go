package main

import "fmt"

type Stringer interface {
	String() string
}

type User struct{ Name string }

func (u User) String() string { return u.Name }

func describe(value any) {
	switch value := value.(type) {
	case nil:
		fmt.Println("nil")
	case Stringer:
		fmt.Println("Stringer:", value.String())
	case int:
		fmt.Println("int:", value)
	default:
		fmt.Printf("unknown: %T\n", value)
	}
}

func main() {
	describe(User{Name: "Go"})
	describe(42)
}
