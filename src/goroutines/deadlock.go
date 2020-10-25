package main

import (
	"sync"
)

var wg sync.WaitGroup

func doStuff(c1, c2 chan int) {
	<-c1    // read from channel 1, wait for the value to e available
	c2 <- 1 // write to channel 2
	wg.Done()
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg.Add(2)
	go doStuff(ch1, ch2) // this goroutine waits to receive something on ch1
	go doStuff(ch2, ch1) // this go routine write to the ch1 but waits to recieve something on ch2
	// above execution casuse a deadlock
	wg.Wait()
}
