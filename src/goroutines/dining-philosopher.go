package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// ChopS ...
type ChopS struct {
	sync.Mutex
}

// Philo ...
type Philo struct {
	leftCS  *ChopS
	rightCS *ChopS
}

func (p Philo) eat(no int) {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("Starting to eat", no)

		p.leftCS.Unlock()
		p.rightCS.Unlock()

		fmt.Println("Finish eating", no)

	}
	wg.Done()
}

func main() {
	const PhilosCount, ChopSticksCount = 5, 5
	var cSticks = make([]*ChopS, ChopSticksCount) // we have {ChopSticksCount} chopsticks
	var philos = make([]*Philo, PhilosCount)      // we have {PhilosCount} philosophers

	for i := 0; i < ChopSticksCount; i++ {
		cSticks[i] = new(ChopS)
	}

	for i := 0; i < PhilosCount; i++ {
		philos[i] = &Philo{cSticks[i], cSticks[(i+1)%ChopSticksCount]}
		wg.Add(1)
	}

	for i := 0; i < PhilosCount; i++ {
		go philos[i].eat(i)
	}

	wg.Wait()
}
