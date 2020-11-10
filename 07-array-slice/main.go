package main

import "fmt"

func main() {
	foo01()
	foo02()
	foo03()
	foo04()

	foo := []int{1, 2, 3, 4}
	calc(1, 2, 3)
	calc(foo...)
	calc(1, 2, 3, 4)
}

func foo01() {
	var foo [3]int
	var bar []int
	test := []int{}

	if bar == nil {
		fmt.Printf("%v\n", foo)
		fmt.Printf("%p\n", &foo)
		fmt.Printf("%p\n", bar)
		fmt.Printf("%p\n", test)
	}
}

func foo02() {
	x := []int{1, 2, 3}
	y := x[1:]
	y[0] = 100
	fmt.Println(x)
	fmt.Println(y)
	y = append(y, 1000)
	y[0] = 99
	fmt.Println(x)
	fmt.Println(y)
}

func foo03() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)
	fmt.Println(x)
}

func foo04() {
	x := []int{1, 2, 3}

	func(arr []int) {
		arr[0] = 7
		fmt.Println(x)
	}(x)
	fmt.Println(x)
}

func calc(vals ...int) {
	foo := 1
	if len(vals) > 0 {
		foo = vals[0]
	}
	fmt.Println(foo)
	fmt.Println("array count:", len(vals))
	for index, val := range vals {
		fmt.Println(index, ":", val)
	}
}
