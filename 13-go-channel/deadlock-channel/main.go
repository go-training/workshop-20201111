package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go func() {
			ch <- 100
			wg.Done()
		}()
	}

	wg.Wait()
}
