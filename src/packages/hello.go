package hello

import "fmt"

var x = 1 // x is not exported
var X = 1 // this X is exported

func printX() { // this is not exported
	fmt.Println(x)
}

func Printx() { // this is exported
	fmt.Println(x)
}
