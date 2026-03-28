package main

import (
	"log"
	"sync"
)

func main() {
	onceFunc := sync.OnceFunc(func() {
		log.Println("only once")
	})

	var wg sync.WaitGroup
	for range 10 {
		wg.Go(func() {
			onceFunc()
		})
	}

	wg.Wait()
}
