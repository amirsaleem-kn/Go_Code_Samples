package main

import (
	"fmt"
	"sync"
)

func printWithWait(value interface{}, wg *sync.WaitGroup) {
	fmt.Println(value)
	wg.Done()
}

/*
 * Main is also a goroutine, each goroutine is executed in a seperate thread (not really :-))
 * We don't know the ordering of go routines execution
 * Some goroutines may not finish as the main may exit after its completion (Main is also a goroutine!)
 */

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go printWithWait("Hello World From First Goroutine", &wg)
	go printWithWait("Hello World From Second Goroutine", &wg)
	go printWithWait("Hello World From Third Goroutine", &wg)
	wg.Wait()
	fmt.Println("Hello World")
}
