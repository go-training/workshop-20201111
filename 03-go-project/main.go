package main

import (
	"fmt"

	"github.com/appleboy/com/random"
)

func hello() string {
	return "歡迎來到 Go 世界"
}

func main() {
	fmt.Println(hello())
	fmt.Println(random.String(10))
}
