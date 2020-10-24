package main

import (
	"fmt"
	"sync"
)

var i int = 0

var mut sync.Mutex

func inc(wg *sync.WaitGroup) {
	mut.Lock()
	i = i + 1
	mut.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go inc(&wg)
	go inc(&wg)
	go inc(&wg)
	go inc(&wg)
	go inc(&wg)
	wg.Wait()

	fmt.Println(i)
}
