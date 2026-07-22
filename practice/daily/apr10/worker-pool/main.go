package main

import (
	"log"
	"sync"
	"time"
)

func work(jobs chan int) {
	for v := range jobs {
		time.Sleep(500 * time.Millisecond)
		log.Println(v)
	}

}
func main() {
	jobs := make(chan int)
	go func() {
		for i := range 20 {
			jobs <- i
		}
		close(jobs)
	}()

	var wg sync.WaitGroup
	poolN := 5

	for range poolN {
		wg.Go(func() {
			work(jobs)
		})
	}
	wg.Wait()
}
