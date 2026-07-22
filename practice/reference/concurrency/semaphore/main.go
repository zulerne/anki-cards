package main

import (
	"fmt"
	"sync"
)

func main() {
	const limit = 2
	sem := make(chan struct{}, limit)
	var wg sync.WaitGroup

	for job := range 5 {
		job := job
		wg.Go(func() {
			sem <- struct{}{}
			defer func() { <-sem }()
			fmt.Println("processed", job)
		})
	}
	wg.Wait()
}
