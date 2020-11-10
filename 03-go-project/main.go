package main

import (
	"fmt"

	"github.com/appleboy/com/random"
)

// Hello ...
func Hello() string {
	return "歡迎來到 Go 世界"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(random.String(10))
}
