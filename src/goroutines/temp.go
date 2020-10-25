package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var no = 0

func g1() {
	no = 4
	wg.Done()
}

func g2() {
	no = no + 1
	wg.Done()
}

func main() {
	wg.Add(2)
	go g1()
	go g2()
	wg.Wait()
	fmt.Println(no)
}
