package main

import (
	"fmt"
)

// By default Array Arguments are COPIED! ( <----- NOTE THIS DUDE )
func sortItUp(arr [3]int) [3]int { // call by value
	return arr // Any update to arr in the function body will not have any effect outside the function
}

// The programmer has to decide when to use call by reference by passing the pointer

// But, this is not the recommended Go way to work with arrays
// you should always use slices, as slices contains pointer to the original array
func sortItUpPointer(arr *[3]int) [3]int { // call by reference
	return *arr // Any update to arr in the function body will reflect in the original array
}

// preferred GO way for pass by reference in GO

func goSortItUp(arr []int) []int {
	arr[0] = 10 // since slice is a pointer, the value passed to this func will contain a pointer
	return arr
}

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(sortItUp(a))
	fmt.Println(sortItUpPointer(&a))

	mySlice := []int{1, 2, 3}
	goSortItUp(mySlice) // use slices, slices are always pass by reference as opposed to array
}
