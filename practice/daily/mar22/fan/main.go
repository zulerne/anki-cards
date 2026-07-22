package main

import (
	"fmt"
	"sync"
)

// Fan-Out: раздаём задачи нескольким воркерам
// Fan-In: собираем результаты в один канал

func fanOut(jobs <-chan int, n int) []<-chan int {
	channels := make([]<-chan int, n)
	for i := range n {
		ch := make(chan int)
		channels[i] = ch
		go func() {
			defer close(ch)
			for j := range jobs {
				ch <- j * 2 // какая-то обработка
			}
		}()
	}
	return channels
}

func fanOut2(jobs <-chan int, n int) []<-chan int {
	channels := make([]<-chan int, n)
	for i := range n {
		ch := make(chan int)
		channels[i] = ch
		go func() {
			defer close(ch)
			for j := range jobs {
				ch <- j * 2
			}
		}()
	}
	return channels
}

func fanIn(channels []<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out) // закрываем когда все воркеры закончили
	}()

	return out
}

func fanIn2(channels []<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for c := range ch {
				out <- c
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	jobs := make(chan int, 10)
	for i := range 10 {
		jobs <- i
	}
	close(jobs)

	results := fanIn2(fanOut2(jobs, 3))
	for r := range results {
		fmt.Println(r)
	}
}
