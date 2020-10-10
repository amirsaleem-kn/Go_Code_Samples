package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// read file loads the entire file into the memory
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	// read stream by stream with the help of the os package
	cursor, e := os.Open("test.txt")

	byteArr := make([]byte, 20)

	cursor.Read(byteArr)

	cursor.Close()

	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(byteArr)

}
