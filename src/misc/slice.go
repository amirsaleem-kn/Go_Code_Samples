package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	s := make([]int, 0, 3)

	var usrInput string

	for usrInput != "X" {
		// accept user input
		fmt.Println("Enter integer value: ")
		fmt.Scan(&usrInput)

		// check if usrInput is an integer or not
		userInput, err := strconv.Atoi(usrInput)

		if err != nil && usrInput != "X" {
			fmt.Println("Please enter integer values")
		} else if usrInput != "X" {
			s = append(s, userInput)
		}
	}

	sort.Ints(s)

	fmt.Println(s)

}
