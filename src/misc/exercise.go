package main

import (
	"fmt"
	"math"
)

// s =½ a t2 + vot + so

// GenDisplaceFn test
func GenDisplaceFn(acceleration, velocity, displacement float64) func(time float64) float64 {
	return func(time float64) float64 {
		return velocity*time + (1 / 2 * (math.Pow((acceleration * time), 2)))
	}
}

func main() {

	var acceleration float64
	var velocity float64
	var displacement float64
	var time float64

	fmt.Println("Enter acceleration: ")
	fmt.Scan(&acceleration)
	fmt.Println("Enter velocity: ")
	fmt.Scan(&velocity)
	fmt.Println("Enter displacement: ")
	fmt.Scan(&displacement)
	fmt.Println("Enter a value for time and the program should compute the displacement: ")
	fmt.Scan(&time)

	fn := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Println("Displacement after 3 seconds is: ")
	fmt.Println(fn(time))
}
