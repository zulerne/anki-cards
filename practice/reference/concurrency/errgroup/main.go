package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	results := make(chan int, 2)

	for _, value := range []int{2, 4} {
		g.Go(func() error {
			select {
			case results <- value * value:
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		})
	}

	if err := g.Wait(); err != nil {
		panic(err)
	}
	close(results)

	for result := range results {
		fmt.Println(result)
	}
}
