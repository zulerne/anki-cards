package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer close(ch)
		for i := range 10 {
			if i == 5 {
				cancel()
				return
			}
			ch <- i
		}
	}()

	var wg sync.WaitGroup
	wg.Go(func() {
		listen(ctx, ch)
	})
	wg.Wait()
}

func listen(ctx context.Context, ch chan int) {
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(v)
		case <-ctx.Done():
			fmt.Println("done")
			return
		}
	}
}
