package main

import "fmt"

/*
 * ---- CALL BY VALUE ----
 * Passed arguments are copied to parameters
 * Modifying parameters has no effect outside the function
 */

/*
 * ---- CALL BY REFERENCE ----
 * Passed arguments are not copied to parameters, instead their pointer is used
 * Modiying parameters updates the original value outside the function
 */

func callByValueSimple(x int) int {
	x = 5
	return x
}

// this function is accepting pointer to the x, instead of the original value x
func callByReferenceSimple(x *int) int {
	*x = 6 // *x means we are updating the data in the memory location
	return *x
}

func main() {
	fmt.Println("Hello World!")

	var x int = 10       // x is assigned the value 10
	callByValueSimple(x) // calling the function which updates the value of x to 5
	fmt.Println(x)       // printing the x will still say 10, the updated value 5 has a new memory location

	callByReferenceSimple(&x) // any update happens to x will happen to the original value of x
	fmt.Println(x)            // x has been updated to 6, since we have passed the pointer to the x
}
