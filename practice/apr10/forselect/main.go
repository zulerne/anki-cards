package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch := make(chan int)
	go func() {
		for i := range 10 {
			time.Sleep(time.Second)
			ch <- i
		}
		close(ch)
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("context done")
			return
		case v, ok := <-ch:
			if ok {
				log.Printf("%v", v)
			} else {
				return
			}
		}
	}
}
