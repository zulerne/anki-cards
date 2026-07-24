package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync.Once — classic Do
	var once sync.Once
	for range 3 {
		once.Do(func() {
			fmt.Println("once.Do: called once")
		})
	}

	// sync.OnceFunc — wraps func() for repeated safe calls
	init := sync.OnceFunc(func() {
		fmt.Println("OnceFunc: initialized")
	})
	init()
	init()

	// sync.OnceValue — wraps func() T, caches the result
	config := sync.OnceValue(func() string {
		fmt.Println("OnceValue: computing")
		return "loaded"
	})
	fmt.Println(config())
	fmt.Println(config())
}
