/*
 * GO doesn't have in built class definition like other languages
 * But you can implement classes in Go
 * See code below
 */

package main

import (
	"fmt"
)

// MyInt type
type MyInt int

// Double method
func (mi MyInt) Double() int {
	return int(mi * 2)
}

// Triple method
func (mi MyInt) Triple() int {
	return int(mi * 3)
}

// ----------- Point Class ----------

// Point struct defines the data variables inside a class
type Point struct {
	x float64
	y float64
}

// Multiply method defines the methods used inside a point class
// We can also specify the reciever as a pointer by passing p *Point, as by default the p is copied
func (p Point) Multiply() float64 {
	return p.x * p.y
}

// Add method
// See, our reciver can be a pointer as well
// And, when we use a pointer in a dot notation, there is no need to
// explicitly deference the points
// so in this case p.x is equivalant to *p.x
func (p *Point) Add() float64 {
	return p.x + p.y
}

// Encapsulation can be achieved by choosing to not to export the private data
// Anything that starts with a capital letter is exported
// Anything that starts with a lowecase letter is not exported to the other files that imports the package

func main() {
	fmt.Println("Hello World")
	v := MyInt(3) // v has a reciever type of int
	v.Double()
	v.Triple()

	// create an object from point class
	p1 := Point{3, 4}

	fmt.Println(p1.Multiply())
}
