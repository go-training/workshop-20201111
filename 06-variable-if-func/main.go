package main

import (
	"errors"
	"fmt"
)

func hello(a, b int) (int, bool, error) {
	if a > b {
		return 1, true, errors.New("Error: a > b")
	}

	return 0, false, errors.New("Error: a < b")
}

func hello2(a, b int) (val int, check bool, err error) {
	if a > b {
		val = 1
		check = true
		err = errors.New("Error: a > b")
		return
	}

	err = errors.New("Error: a < b")
	return
}

func variable() {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)
	fmt.Println()
}

func main() {
	_, check, _ := hello(1, 2)
	if check {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if _, _, err := hello(1, 2); err != nil {
		fmt.Println("true")
	} else {
		fmt.Println(err)
	}

	if _, _, err := hello(1, 2); err == nil {
		fmt.Println("true")
	} else {
		fmt.Println(err)
	}

	variable()

	for i := 0; i < 5; i++ {
		defer fmt.Println("defer to:", i)
		defer func() {
			fmt.Println("defer func to:", i)
		}()
	}

	var foo = map[string]string{
		"foo": "test01",
		"bar": "test02",
	}

	for i, v := range foo {
		fmt.Println(i, v)
	}

	c := make(chan int, 10)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	for v := range c {
		fmt.Println("chan value:", v)
	}
}
