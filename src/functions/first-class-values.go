package main

import (
	"fmt"
)

/*
 * -------------------------------------------------
 *    Functions can be treated like other types
 * -------------------------------------------------
 * -> Variables can be declared with a function type
 * -> Function can be created dynamically at run time insider another functions
 * -> Functions can be passed as argument
 * -> Functions can be returned as values
 * -> Functions can be stored in data structures like slices struct etc
 */

// returning a function
func product(x int) func(y int) int {
	return (func(y int) int {
		return x * y
	})
}

// passing a function as an argument
func applyIt(f func(int) int, value int) int {
	return f(value)
}

func increment(value int) int {
	return value + 1
}

// anonymous function
// Function without names are called anonymouse functions

func main() {
	fmt.Println("Hello World")
	multiplyBy2 := product(2)
	multiplyBy4 := product(4)
	fmt.Println(multiplyBy2(2))
	fmt.Println(multiplyBy4(4))
	fmt.Println(applyIt(func(value int) int { return value + 1 }, 10)) // anonymous function
	fmt.Println(applyIt(increment, 10))                                // non-anonymous function
}
