/*
 * ------------ Structs in GO -------------
 * Struct groups together objects of arbitary data type
 *
 * Example: Person Struct with name, phone, and an address
 */

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name  string
	addr  string
	phone string
}

func main() {

	var amir person
	var amirCopy person

	amir.name = "Amir Saleem"
	amir.addr = "India"
	amir.phone = "9582678827"

	fmt.Println(amir)

	barr, err := json.Marshal(amir)

	fmt.Println(barr)

	if err != nil {
		fmt.Println(err)
	}

	unmarshallErr := json.Unmarshal(barr, &amirCopy)

	if unmarshallErr != nil {
		fmt.Println(unmarshallErr)
	}

	fmt.Println(amirCopy)

}
