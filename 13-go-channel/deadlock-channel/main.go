package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
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
