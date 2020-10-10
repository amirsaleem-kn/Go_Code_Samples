/**
 * ----------- Constants in GO ------------
 */

package main

import "fmt"

func main() {
	const x = 10 // x is a constant, it cannot be reassigend. Constant must be declared and assigned at the same time
	const (
		z = 1
		k = 2
	)

	const (
		monday = iota
		tuesday
		wednesday
	)

	if 1 == 1 {
		fmt.Println("Yes, 1 equals to 1")
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// in go there is no need to put break statements, the control automatically breaks. You can put break thought but that is not needed
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	case 3:
		fmt.Println("x is 3")
	}

	// tagless switch as an alternative to if-else-if-else conditions

	switch {
	case x == 0:
		fmt.Println("x is 0")
	case x > 1:
		fmt.Println("x is greater than 1")
	}

	// reads user input
	var name string = ""
	fmt.Scan(&name)
	fmt.Println(name)
}
