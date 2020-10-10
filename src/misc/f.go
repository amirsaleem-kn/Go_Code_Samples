// read project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	type Name struct {
		fname string
		lname string
	}
	names := []Name{}

	fmt.Println("Enter file name: ")
	fileName := ""
	fmt.Scan(&fileName)

	barr, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fstrings := strings.Split(string(barr), "\n")
	for _, s := range fstrings {
		tmp := strings.Split(s, " ")
		names = append(names, Name{fname: tmp[0], lname: tmp[1]})
	}

	for _, n := range names {
		fmt.Printf("fname: %s lname: %s\n", n.fname, n.lname)
	}
}
