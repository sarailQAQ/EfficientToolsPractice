package main

import (
	"fmt"
	"os"
)

func main() {
	str := os.Getenv("message")
	if len(str) == 0 {
		str = "hello dear friend!"
	}
	fmt.Println(str)
}
