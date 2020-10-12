package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

// if all arguments are same, we can skip the data type but the last

func product(x, y, z int) int {
	return x * y * z
}

// functions can return more than one values

func giveMeBackMyArgs(a, b, c int) (int, int, int) {
	return a, b, c
}

func printHello() {
	fmt.Println("The sum is:", sum(4, 4))
	fmt.Println("The product is:", product(2, 2, 2))
	a, b, c := giveMeBackMyArgs(2, 2, 2)
	fmt.Println(a, b, c)
}

func main() {
	printHello()
}
