package main

import (
	"fmt"
	"time"
)

func process(i int) {
	fmt.Println(i)
}

func main() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	requests := make(chan int)
	go func() {
		for i := range 10 {
			requests <- i
		}
		close(requests)
	}()

	for t := range requests {
		<-ticker.C
		process(t)
	}
}
