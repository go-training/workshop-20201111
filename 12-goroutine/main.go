package main

import (
	"fmt"
	"sync"
	"time"
)

func foo(input string, wg *sync.WaitGroup) {
	fmt.Println(input)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	// one
	go foo("1", wg)
	go foo("2", wg)
	wg.Wait()

	// two
	msg := "let's go"
	go func(m string) {
		fmt.Println(m)
	}(msg)
	msg = "let's gogogo"

	go waitGroup()
	time.Sleep(2 * time.Second)
}

func waitGroup() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			// defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
			time.Sleep(500 * time.Millisecond)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("m:", len(m))
	fmt.Println("m:", m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
