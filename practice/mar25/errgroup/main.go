package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

func work(ctx context.Context, i int) error {
	if i == 5 {
		return fmt.Errorf("work stopped")
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		fmt.Println(i)
		return nil
	}
}
func main() {
	g, ctx := errgroup.WithContext(context.Background())

	for n := range 10 {
		g.Go(func() error {
			return work(ctx, n)
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Print("no errors found")
}
