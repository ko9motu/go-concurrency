package clonerequest

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork(
	done <-chan interface{},
	id int,
	wg *sync.WaitGroup,
	result chan<- int,
) {
	started := time.Now()
	defer wg.Done()

	// Simulate random load
	simulatedLoadtime := time.Duration(1+rand.Intn(5)) * time.Second
	select {
	case <-done:
	case <-time.After(simulatedLoadtime):
	}

	select {
	case <-done:
	case result <- id:
	}

	took := time.Since(started)

	// Display how long handlers would have taken
	if took < simulatedLoadtime {
		took = simulatedLoadtime
	}
	fmt.Printf("%v took %v\n", id, took)
}
