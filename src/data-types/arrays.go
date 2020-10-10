/**
 * ------------ Arrays in GO -----------
 * Arrays in GO works as the traditional definition of array says. Like, array is a data-type
 * of fixed length which contains homogeneous data type elements
 */

package main

import "fmt"

func main() {
	// Array
	var x [5]int
	x[0] = 2
	fmt.Println(x)

	// Array Literal -> Predefined values
	// ... referes to the size of the array literal, we can specify the size with an integer or it can infer the same using ...
	var y = [...]int{1, 2, 3, 4, 5} // ... to infer the size of the array literal
	var z = [5]int{1, 2, 3, 4, 5}   // integer to determine the size of the array literal
	fmt.Println(y)

	// iterating an Array
	for v := range x {
		fmt.Println(v)
	}

	// iterating an Array
	for i, v := range z {
		fmt.Println(v, i)
	}
}
