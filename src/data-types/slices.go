/*
 * ------- Slices in GO -----
 * Unlike arrays, slices don't have fixed length, they have dynamic length
 * A slice is just a subset of an original array
 *
 * Technically, a slice is just a window on an underlying array. You can imagine slice as a
 * sub-array from a arge array, so if an array has a size 0 - n, the slice can be a subset of an array like
 * 5 - n-5
 *
 * Properties of a slice
 * 1. Pointer indicates the start of the slice -> which means every slice has a start pointer which is a pointer to an original array
 * 2. Length -> This indicates the number of elements in the slice
 * 3. Capacity -> This defines the maximum number of elements that can be inserted into a slice
 *
 * Note, any update to the slice happens to the original array, so if you update an element of a slice, you are updating
 * the original array element, and all the slices which refer to the main array will be updated
 */

package main

import (
	"fmt"
)

func main() {
	// This is our main array, this has a length of 8
	mainArr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}

	// lets create a slice out of the main array
	slice1 := mainArr[1:3] // first element is the pointer to the array first element, the second element is the index of the last element + 1
	slice2 := mainArr[2:5]

	// ....Since slice is just a pointer to the main array
	// Any update to the original array will be reflected in the slice1 and slice2 as well
	// Any update to the slice1 will be reflected in the slice2 and mainArr as well
	// Any update to the slice2 will be reflected in the slice1 and mainArr as well

	fmt.Println(slice1) // [b c]
	fmt.Println(slice2) // [c d e]

	fmt.Println("Length of slice1", len(slice1))   // 2
	fmt.Println("Capacity of slice1", cap(slice1)) // 7 because we are starting from index 1 and not 0 hence cap is 7 and not 8

	fmt.Println("Length of slice2", len(slice2))   // 3
	fmt.Println("Capacity of slice2", cap(slice2)) // 6 because we are starting from index 2 and not 0 hence cap is 6 and not 8

	/* ----------- Slice Literals ------------ */

	// Just like array literals, we have slice literals as well
	// the slice literals creates an underlying array, and then a slice to refer to it
	// one thing to note here is that, in the slice literals the len and cap of the slice are same

	sliceLiteral := []int{1, 2, 3} // since we have not defined the length using integer or ..., this is a slice literal

	/* ----------- Variable Slices ----------- */
	// create a slice (and array) using make method

	s := make([]int, 10) // create an integer slice of length 10 and capacity 10

	k := make([]int, 10, 100) // create an integer slice of length 10 and capacity 100

	/* -------- Adding a new element to the slice with an append function ------- */
	// The append function keeps on increasing the size of the array, so you never run in an overflow situation when using append
	k = append(k, 5)

}
