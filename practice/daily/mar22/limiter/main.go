package main

import (
	"fmt"
	"time"
)

func process(i int) {
	fmt.Printf("%v\n", i)
}
func main() {
	limiter := time.NewTicker(100 * time.Millisecond)
	defer limiter.Stop()

	requests := make(chan int, 10)
	for n := range 10 {
		requests <- n
	}
	close(requests)

	for r := range requests {
		<-limiter.C
		process(r)
	}
}
