/**
 * ------------ Strings in Go -----------
 */

/**
 * In Go, or generally most of the programming languages, strings are represented as a sequence of bytes.
 * There are two types of standards we use in defining strings
 * ASCII -> American Standard Code for Information Exchange: 8 bit long code, maximum 256 possible characters can be represented,like A-Z, a-z, 1-10
 * Unicode -> Character encoding, each character is associated with an (7)8-bit number, 32 bit long code as opposed to 8 bits in ASCII, this gives us
 * flexibility in representing characters other than English like chinese, and local language symbols and characters
 * UTF-8 code are same as ASCII as both of them uses 8 bits,
 * In GO, these character codes in Hexadecimal are called Runes, A is 41 rune character
 *
 * Strings in GO are read-only, when updated a new memory location is allocated
 * Strings are immutable
 */

package main

import (
	"fmt"
)

// func commonStrFunctions() {
// 	const x string = "Hello"
// 	isDigit(x)
// 	isSpace()
// 	isLetter()
// 	x.ToUpper()
// 	x.ToLower()
// }

func main() {
	const y string = "Hello there!"
	x := "My String!"
	fmt.Println(x, y)
}
