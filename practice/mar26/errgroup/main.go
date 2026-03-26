package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

func work(ctx context.Context, i int) error {
	if i == 5 {
		return fmt.Errorf("work failed")
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		log.Printf("work %d", i)
		return nil
	}
}
func main() {
	g, ctx := errgroup.WithContext(context.Background())

	for i := range 10 {
		g.Go(func() error {
			return work(ctx, i)
		})
	}
	if err := g.Wait(); err != nil {
		log.Fatalf("group error %v", err)
	}
	log.Println("group done")
}
