package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

func work(ctx context.Context, n int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if n == 5 {
		return fmt.Errorf("work failed")
	}
	log.Println(n)
	return nil
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	for i := range 10 {
		g.Go(func() error {
			return work(ctx, i)
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Println("work done")
}
