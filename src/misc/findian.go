package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter test string: ")
	name, err := in.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	lowerCaseName := strings.ToLower(strings.TrimSuffix(name, "\n"))

	if strings.HasPrefix(lowerCaseName, "i") && strings.HasSuffix(lowerCaseName, "n") && strings.Contains(lowerCaseName, "a") {
		fmt.Println("FOUND")
	} else {
		fmt.Println("NOT FOUND")
	}
}
