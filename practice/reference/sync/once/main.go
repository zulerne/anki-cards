package main

import (
	"fmt"
	"sync"
)

var config = sync.OnceValue(func() string {
	return "loaded once"
})

func main() {
	fmt.Println(config())
	fmt.Println(config())
}
