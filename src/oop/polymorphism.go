/*
 * Polymorphism is a property which defines an object can have different form
 * or a method can have different forms
 */

package main

import (
	"fmt"
)

// Shape2D Defining an interface
type Shape2D interface {
	Area() float64
	Perimeter() float64
}

func main() {
	fmt.Println("Hello World")
}
