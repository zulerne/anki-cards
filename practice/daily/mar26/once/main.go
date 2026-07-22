package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	// var once sync.Once
	var wg sync.WaitGroup

	once := sync.OnceFunc(func() {
		log.Println("once")
	})
	onceValue := sync.OnceValue(func() int {
		return 10
	})
	for i := range 10 {
		wg.Go(func() {
			// once.Do(func() {
			// 	log.Println("once")
			// })
			once()
			val := onceValue()
			fmt.Printf("val: %v\n", val)
			log.Print(i)
		})
	}
	wg.Wait()
}
