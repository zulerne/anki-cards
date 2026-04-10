package main

import (
	"log"
	"time"
)

func main() {
	requests := make(chan int, 10)
	for i := range 10 {
		requests <- i
	}
	close(requests)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for r := range requests {
		<-ticker.C
		log.Printf("req %v:", r)
	}
}
