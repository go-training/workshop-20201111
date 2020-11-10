package main

import (
	"fmt"

	"foobar/hello"
	"foobar/hello2"
	"foobar/hello3"
)

func main() {
	fmt.Println(hello.Hello())
	fmt.Println(hello2.Hello())
	fmt.Println(hello3.Hello())
}
