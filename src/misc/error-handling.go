package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("some/file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
	fmt.Println("Hello")
}
