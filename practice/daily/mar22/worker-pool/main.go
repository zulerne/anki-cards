package main

import (
	"fmt"
	"sync"
)

func work(tasks chan int) {
	for t := range tasks {
		fmt.Printf("%v\n", t)
	}
}
func main() {
	poolN := 10
	jobs := make(chan int)
	var wg sync.WaitGroup

	for range poolN {
		wg.Go(func() {
			work(jobs)
		})
	}

	for _, v := range []int{1, 3, 6, 2, 6, 1, 7} {
		jobs <- v
	}
	close(jobs)
	wg.Wait()
}
