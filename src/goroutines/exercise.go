package main

import (
	"fmt"
	"sort"
)

func sortInts(arr []int, ch chan []int) {
	sort.Ints(arr)
	ch <- arr
}

func main() {
	ch := make(chan []int)
	inputArr := [...]int{5, 6, 3, 1, 0, -1, 4, -8}

	go sortInts(inputArr[0:2], ch)
	go sortInts(inputArr[2:4], ch)
	go sortInts(inputArr[4:6], ch)
	go sortInts(inputArr[6:8], ch)

	first := <-ch
	second := <-ch
	third := <-ch
	fourth := <-ch

	s := make([]int, 0, 10)

	s = append(s, first...)
	s = append(s, second...)
	s = append(s, third...)
	s = append(s, fourth...)

	sort.Ints(s)

	fmt.Println(s)
}
