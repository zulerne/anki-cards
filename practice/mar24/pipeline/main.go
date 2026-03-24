package main

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range 10 {
			select {
			case out <- i:
				time.Sleep(300 * time.Millisecond)
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func square(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range ch {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func print(ch <-chan int) {
	for v := range ch {
		fmt.Printf(" %v\n", v)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	print(square(gen(ctx)))
}
