package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	semaphor := make(chan struct{}, 10)

	for j := range 20 {
		wg.Go(func() {
			defer func() { <-semaphor }()
			semaphor <- struct{}{}
			time.Sleep(time.Second)
			fmt.Println(j)
		})
	}
	wg.Wait()
}
