package main

import (
	"fmt"
	"sync"
)

func fanIn(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Go(func() {
			for value := range input {
				out <- value
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func values(values ...int) <-chan int {
	out := make(chan int, len(values))
	for _, value := range values {
		out <- value
	}
	close(out)
	return out
}

func main() {
	for value := range fanIn(values(1, 3), values(2, 4)) {
		fmt.Println(value)
	}
}
