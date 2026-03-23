package main

import (
	"fmt"
	"sync"
)

func fanOut(in <-chan int, n int) []<-chan int {
	channels := make([]<-chan int, n)

	for i := range n {
		ch := make(chan int)
		channels[i] = ch
		go func() {
			defer close(ch)
			for j := range in {
				ch <- j
			}
		}()
	}
	return channels
}

func fanIn(channels []<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(ch)
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
	for i := range out {
		fmt.Println(i)
	}
}
