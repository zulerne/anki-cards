package main

import (
	"fmt"
	"sync"
)

func work(jobs chan int) {
	for v := range jobs {
		fmt.Printf("v: %v\n", v)
	}
}
func main() {
	jobs := make(chan int)

	poolN := 5
	var wg sync.WaitGroup

	for range poolN {
		wg.Go(func() {
			work(jobs)
		})
	}

	for i := range 10 {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
