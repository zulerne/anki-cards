package main

import (
	"fmt"
	"sync"
	"time"
)

func genJobs(n int) <-chan int {
	jobs := make(chan int)
	go func() {
		for i := range n {
			jobs <- i
		}
		close(jobs)
	}()
	return jobs
}

func main() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 5)

	for v := range genJobs(20) {
		wg.Go(func() {
			defer func() { <-sem }()
			sem <- struct{}{}
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("v: %v\n", v)
		})
	}

	wg.Wait()
}
