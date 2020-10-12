package main

import "fmt"

/**
 * ------------- How to write good GO functions that everyone loves :) -------------
 * 1. Orgnaisation is the key, organise your functions properly
 * 2. Give meaningful names to your functions
 * 3. Don't give too long names to your function, but if need to be, you can go for it
 * 4. Always try to follow Function Coheision Principle -> Your function should do only single operation ( there are cases in this :-) )
 * 5. Limit the number of parameters -> ideal is 0-2. If required you can use struct, slices and maps to avoid large number of arguments
 * 6. Don't add too much code in a single function - Fat functions are unfit :)
 */

// bad appraoch (Too many arguments!)
func computeTriangleAreaBad(x1, y1, z1, x2, y2, z2, x3, y3, z3 int) int {
	return 10 // just a random number
}

// better / good / nice approach with struct

type point struct{ x, y, z float32 }

func computeTriangleAreaGood(p1, p2, p3 point) int {
	return 10
}

// ideal approach
type trinangle struct{ p1, p2, p3 point }

func computeTriangleAreaIdeal(t trinangle) int {
	return 10
}

func main() {
	fmt.Println("Hello")
}
