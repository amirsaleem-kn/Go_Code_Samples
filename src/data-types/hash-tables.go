/**
 * ------------ Hash Tables in GO -----------
 * A hash table contains a key value pair
 * Each value is associated with a unique key
 * A hash function is used to compute the slot for a key
 *
 * In GO Languages, a map is an implementation of a Hash Table
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Test")
	var idMap map[string]int
	idMap = make(map[string]int)

	fmt.Println(idMap)

	// map literal
	var c = map[string]int{
		"amir":  1,
		"gulam": 2}

	fmt.Println(c)

	// accesing a key
	fmt.Println(c["amir"])

	// adding a new key
	c["john"] = 454

	// deleting a key
	delete(c, "john")

	// length of a map
	fmt.Println(len(c))

	// two value assignment, value, ifAlreadyExsists | bool := c["amir"]
	amirID, isExists := c["amir"] // returns id, and a boolean to determine if key exists or not
	fmt.Println(amirID, isExists)

	// iterating through a map
	for key, value := range c {
		fmt.Println(key, value)
	}

}
