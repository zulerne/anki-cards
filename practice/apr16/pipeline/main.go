package main

import (
	"context"
	"fmt"
	"time"
)

func generate(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range 10 {
			select {
			case ch <- i:
				time.Sleep(200 * time.Millisecond)
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}

func print(ch <-chan int) {
	for v := range ch {
		fmt.Printf("v: %v\n", v)
	}
}

func double(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range ch {
			out <- v * v
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	print(double((generate(ctx))))
}
