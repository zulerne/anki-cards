package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for value := range 3 {
		value := value
		wg.Go(func() {
			fmt.Println(value)
		})
	}
	wg.Wait()
}
