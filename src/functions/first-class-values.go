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

// anonymous function
// Function without names are called anonymouse functions

func main() {
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
}
