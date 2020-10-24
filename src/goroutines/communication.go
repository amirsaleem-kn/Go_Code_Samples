package main

import (
	"fmt"
	"sync"
)

func multiply(a int, b int, wg *sync.WaitGroup, ch chan int) {
	result := a * b
	ch <- result
	ch <- 3
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	const a = 3
	const b = 4
	const c = 2
	const d = 2

	wg.Add(2)

	ch := make(chan int)

	go multiply(a, b, &wg, ch)
	go multiply(c, d, &wg, ch)

	first := <-ch
	second := <-ch

	wg.Wait()

	fmt.Println(first, second)
}
