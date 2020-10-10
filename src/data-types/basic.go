package main

import (
	"fmt"
)

// Integers

func typesOfint() {
	var x int
	var a int8
	var b int16
	var c int32
	var d int64
	var e uint8
	var f uint16
	var g uint32
	var h uint64
}

func typesOfFloats() {
	var c float32
	var d float64
	var k complex64 = complex(2, 3) // 2 real and 3 imaginary 2+3i
}

// type conversion

func failOnDiffTypes() {
	var x int32 = 1
	var y int16 = 2
	// x = y // This will fail as the compiler sees these two numbers as different types, and handles them differently.
	// to solve the above problem we first need to make their types equal
	x = int32(y) // this will work, as we have converted y to int32
}

func main() {
	fmt.Println("Main Executed")
}
