package main

import (
	"context"
	"fmt"
)

func generate(ctx context.Context, values ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, value := range values {
			select {
			case out <- value:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func square(ctx context.Context, input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range input {
			select {
			case out <- value * value:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for value := range square(ctx, generate(ctx, 1, 2, 3, 4)) {
		fmt.Println(value)
	}
}
