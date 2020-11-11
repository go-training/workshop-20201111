package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	go func(ch <-chan int) {
		wg.Add(1)
		fmt.Println(<-ch)
		wg.Done()
	}(ch)
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go func(ch chan<- int) {
			ch <- 100
			wg.Done()
		}(ch)
	}

	wg.Wait()
}
