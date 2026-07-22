package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	done := make(chan int)
	var wg sync.WaitGroup

	wg.Add(10)
	go func() {
		for i := range 10 {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for {
			select {
			case val := <-ch:
				wg.Done()
				fmt.Println(val)
			case <-done:
				return
			}
		}
	}()
	wg.Wait()
}
