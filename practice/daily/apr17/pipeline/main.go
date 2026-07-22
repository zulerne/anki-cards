package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range n {
			select {
			case <-ctx.Done():
				return
			case out <- v:
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()
	return out
}

func print(in <-chan int) {
	for v := range in {
		fmt.Printf("v: %v\n", v)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	print(square(gen(ctx, 100)))
}
