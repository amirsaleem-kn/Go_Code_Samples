package main

import "fmt"

// Speaker test
type Speaker interface { //  interface
	Speak()
}

// Shape2D ..
type Shape2D interface {
	Area() float64
	Parameter() float64
}

/*
 * Where to use interface
 * 1. Need a function which takes multiple types of parameters
 * 2. Need objects which implements the same methods
 */

// Dog test
type Dog struct {
	name string
}

// Cat test
type Cat struct {
	name string
}

// Speak ..
func (d Dog) Speak() { // method Speak which takes Dog as an argument
	fmt.Println(d.name)
}

// Speak Cat
func (c Cat) Speak() {
	fmt.Println(c.name) // Speak for cat
}

// Empty Interface
// With empty interface, the function can accept any type of arguments
func printMe(val interface{}) {
	fmt.Println(val)
}

func main() {
	var s1 Speaker        // Concrete Type
	var d1 = Dog{"Brian"} // Dynamic Type and Dynamic Value
	s1 = d1
	s1.Speak() // here the compiler knows that we want to call Speak method of Dog

	var c1 = Cat{"Kitty"}
	c1.Speak() // the compiler knows that we want to call Speak method of cat

	printMe("Hi")

	cat, ok := s1.(Cat) // type assertions
	fmt.Println(cat, ok)

	// swtich based type assertions
	switch sh := s1.(type) {
	case Cat:
		fmt.Println("Its a cat", sh)
	case Dog:
		fmt.Println("Its a dog", sh)
	}
}
