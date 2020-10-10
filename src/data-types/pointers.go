/**
 * ------------------------ POINTERS IN GO ------------------------
 */

/** A pointer is an address to a data in memory. For e.g if a variable X has a value 100 stored in memory location 0xc00008a008
 * Then 0xc00008a008 is the memory address of X which is 100. The pointer p holds this memory address instead of the actual value
 * so if we try to access the data stored in location 0xc00008a008 using a pointer we will get 10
 *
 * A pointer has two components. &p and *p. If you access pointer p with the ampersand operator, you will get the value of address
 * the pointer p is pointing to like &p and if you access p with the star operator, you will get the actual value the pointer p
 * is references to like *p
 */

package main

import "fmt"

func main() {

	var x int = 1
	var y int

	fmt.Println("x is: ", x)
	fmt.Println("y is: ", y)

	var ip *int // ip is pointer to int which means ip will point to integer values upon initialization

	ip = &x // ip is now pointing to variable x's memory address ( Virtual Memory Address! )
	y = *ip // y is now 1, since start operator gives us the value stored in that location where ip is ponting to (Dereferencing)

	fmt.Println("*ip is", *ip) // actual value stored in a memory address stored in ip
	fmt.Println("&ip is", &ip) // memory address of x

	// The new() function is another way to create variables, the new() function declares the variable and returns the memory address aka pointer
	var p = new(int) // p is a pointer to int
	*p = 3           // the value stored in p's location is 3 and &p will give us the memory address of 3

	fmt.Println("p is a pointer to int: ", p)
	fmt.Println("*p is a value stored in the location of p: ", *p)

}
