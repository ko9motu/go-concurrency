package channel

import (
	"fmt"
	"time"
)

func SelectExample1() {
	var c1, c2 <-chan interface{}
	var c3 chan<- interface{}

	select {
	case <-c1:
		// do something
	case <-c2:
		// do something
	case c3 <- struct{}{}:
		// do something
	}
}

func SelectExample2() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Broking on read...")

	select {
	case <-c:
		fmt.Printf("Unbroking %v later.\n", time.Since(start))
	}
}

func SelectExample3() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 999; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
	fmt.Printf("c1count: %d\nc2count: %d\n", c1Count, c2Count)
}

func SelectExample4() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:

		}
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achived %v cycles of work before signalled to stop.\n", workCounter)
}
