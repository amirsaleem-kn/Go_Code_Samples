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
// Note: functions uses a Lexical Scoping. so all the variables of a top level functions are available in the inner functions as well
func multiply(x int) func(y int) int {
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

// Sample closure in Go
func closure() func() int {
	var x int = 0
	return func() int { // this function has access to x defined in the parent environment
		x++
		return x
	}
}

// Variable functions in Go
func varArgs(args ...int) int {
	// args is treated as a slice inside a function
	return args[0]
}

// Deferred function calls
// Deferred functions are called when all the surrounding functions gets called
// Helpful in writing a cleanup code, like closing a file, closing a db connection etc
// Gets called in the end of a function, just put a defer keyword prefix with a function call
func iamADeferredFunc() {
	fmt.Println("Bye! from a deferred function")
}

// anonymous function
// Function without names are called anonymouse functions

func main() {
	// Although this is called in the beginning, since we have prexied the function call with defer keyword
	// this function will be called when all the existing code in main is called
	defer iamADeferredFunc()
	fmt.Println("Hello World")
	multiplyBy2 := multiply(2) // this returns a new function
	multiplyBy4 := multiply(4) // this also returns a new function
	fmt.Println(multiplyBy2(2))
	fmt.Println(multiplyBy4(4))
	fmt.Println(applyIt(func(value int) int { return value + 1 }, 10)) // anonymous function
	fmt.Println(applyIt(increment, 10))                                // non-anonymous function
	// closure execution
	inc := closure()   // first time when we execute this function, it creates an env for x and returns a function
	fmt.Println(inc()) // now each time we call inc, the value of x is updated in the parent function
	fmt.Println(inc()) // so value of x is incremented each time we call inc method
	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(varArgs(1, 2, 3, 4))
	fmt.Println(varArgs([]int{1, 2, 3, 4, 5}...)) // slices can also be expanded to a variable number of arguments
}
