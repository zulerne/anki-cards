package main

import (
	"fmt"
	"sync"
)

func fanOut(in <-chan int, n int) []<-chan int {
	out := make([]<-chan int, n)

	for i := range n {
		ch := make(chan int)
		out[i] = ch
		go func() {
			defer close(ch)

			for v := range in {
				ch <- v
			}
		}()
	}

	return out
}

func fanIn(in []<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range in {
		wg.Go(func() {
			for v := range ch {
				out <- v
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch := make(chan int, 10)
	for i := range 10 {
		ch <- i
	}
	close(ch)

	out := fanIn(fanOut(ch, 3))

	for v := range out {
		fmt.Printf("v: %v\n", v)
	}
}
