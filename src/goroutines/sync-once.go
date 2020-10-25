package main

import (
	"fmt"
	"sync"
)

var on sync.Once

func setup() {
	fmt.Println("Init")
}

func doStuff(wg *sync.WaitGroup) {
	on.Do(setup)
	fmt.Println("Hello")
	wg.Done()
}

func doAnotherStuff(wg *sync.WaitGroup) {
	on.Do(setup)
	fmt.Println("Hello another stuff")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go doStuff(&wg)
	go doStuff(&wg)
	go doAnotherStuff(&wg)
	wg.Wait()
}
