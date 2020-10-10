/*
 * --------- Variable Scope in GO --------
 */

/*
 * In GO, there are two scopes, local and global. The global scope is available to all the functions in a file, while
 * a local scope is local to the block and is not accessible outside of the block
 *
 * a block is identified anything in between curly ( {} ) braces. So functions, if statements and for loops have their own
 * block and their own scope of variables
 * {
 *   {
 *     {
 *         {} Function Block and if, switch, for loop blocks inside a file block
 *     } // File block
 *   } // package block contains all packages
 * } // universe block contains all Go source code
 *
 * Go is a Lexically Scoped language, which means that the program recursively tries to find the variable inside a block until
 * it reaches the universe block
 */

package main

import (
	"fmt"
)

var x = 1 // This variable has a global scope and can be accessed anywhere in this file

func printGlobalX() {
	// since x is globally defined, this function has access to the variable x
	fmt.Println("x is: ", x)
}

func printLocalX() {
	// if we re-define the same var x in this func, the function will reference to the local variable instead
	// This will not throw error
	var x = 2 // local variable
	fmt.Println("local x is: ", x)
}

func checkIfBlock() {
	if true {
		var x = 100 // this variable is available only in the if block
		fmt.Println("x in if block: ", x)
	}

	fmt.Println("x outside of block: ", x)
}

func main() {
	printGlobalX()
	printLocalX()
	checkIfBlock()
}
