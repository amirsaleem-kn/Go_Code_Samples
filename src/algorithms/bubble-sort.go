package main

import (
	"fmt"
	"strconv"
)

func printDirections() {
	fmt.Println()
	fmt.Println("************* GO Program to Sort the Slice using BubbleSort *************")
	fmt.Println("------> Please enter upto 10 integers")
	fmt.Println("------> Press X to exit")
	fmt.Println("------> Enter digits in new line")
	fmt.Println("*************************************************************************")
	fmt.Println()
}

// swap swaps two numbers
// It returns the two values which is the swapped version of the input
func swap(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

// BubbleSort sorts an array using bubble sort algorithm
func BubbleSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = swap(arr[i], arr[j])
			}
		}
	}

}

func main() {
	sl := []int{}

	printDirections()

	var usrInput string

	for usrInput != "X" && len(sl) < 10 {
		// accept user input
		fmt.Scan(&usrInput)

		// convert the user input value to integer value
		userInput, err := strconv.Atoi(usrInput)

		if err != nil && usrInput != "X" {
			fmt.Println("Please enter integer values")
		} else if usrInput != "X" {
			sl = append(sl, userInput)
		}
	}

	fmt.Println("Sorted Slice:")
	BubbleSort(sl)

	fmt.Println(sl)

}
