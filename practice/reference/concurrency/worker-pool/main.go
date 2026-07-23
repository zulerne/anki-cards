package main

import (
	"context"
	"fmt"
	"sync"
)

func worker(ctx context.Context, jobs <-chan int, results chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			select {
			case results <- job * job:
			case <-ctx.Done():
				return
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	jobs := make(chan int)
	results := make(chan int)
	var wg sync.WaitGroup

	for range 2 {
		wg.Go(func() { worker(ctx, jobs, results) })
	}
	go func() {
		defer close(jobs)
		for _, job := range []int{1, 2, 3, 4} {
			jobs <- job
		}
	}()
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
