package main

import (
	"fmt"
	"sync"
)

func work(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	poolN := 5
	jobs := make(chan int)
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
