package main

import "fmt"

/*
Race condition occurs when two programs depends upon the same memory, and since execution order of goroutines is
non-deterministic, instructions within the goroutines are also non-deterministic, we might get different values of the shared data
based on the execution order of the tasks done in tasks interleaving.
*/

// This is a global variable
var count int

// Method to increment the counter value
func updateCount() {
	count++
}

// Method to read the counter value
func readCount() int {
	fmt.Println(count)
	return count
}

func main() {
	go updateCount()   // update the count in a goroutine
	go readCount()     // read the count in a go routine
	fmt.Println(count) // print the counter variable
}
