package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go func(ch <-chan int) {
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch)

	wg.Add(1)
	go func(ch chan<- int) {
		for i := 1; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
