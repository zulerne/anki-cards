package main

import (
	"fmt"
	"sync"
	"time"
)

func work() {
	time.Sleep(time.Second)
	fmt.Println("work done")
}
func main() {
	sem := make(chan struct{}, 10)
	var wg sync.WaitGroup

	for range 20 {
		wg.Go(func() {
			defer func() { <-sem }()
			sem <- struct{}{}
			work()
		})
	}
	wg.Wait()

}
