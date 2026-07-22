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

func run(ctx context.Context, jobs []int, workerCount int) []int {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	jobCh := make(chan int)
	resultCh := make(chan int)
	var wg sync.WaitGroup

	for range workerCount {
		wg.Go(func() {
			worker(ctx, jobCh, resultCh)
		})
	}

	go func() {
		defer close(jobCh)
		for _, job := range jobs {
			select {
			case jobCh <- job:
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var results []int
	for result := range resultCh {
		results = append(results, result)
	}
	return results
}

func main() {
	results := run(context.Background(), []int{1, 2, 3, 4}, 2)
	fmt.Println(results)
}
