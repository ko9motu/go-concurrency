package concurrencypattern

import (
	"fmt"
	"math/rand"
	"time"
	// "crypto/rand"
)

// AP = AntiPattern
func APGorutineWriteChannel() {
	newRandCh := func() <-chan int {
		randCh := make(chan int)

		go func() {
			defer fmt.Println("newRandCh closure exited.") // <- never do
			defer close(randCh)

			for {
				randCh <- rand.Int()
			}
		}()

		return randCh
	}

	randCh := newRandCh()
	fmt.Println("3 randome ints")

	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randCh)
	}
}

func GorotineWriteChannel() {
	newRandCh := func(done <-chan interface{}) <-chan int {
		randCh := make(chan int)
		go func() {
			defer fmt.Println("newRandCh closure exited.")
			defer close(randCh)

			for {
				select {
				case randCh <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randCh
	}
	done := make(chan interface{})
	randCh := newRandCh(done)

	fmt.Println("3 rondom ints")

	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randCh)
	}
	close(done)
	time.Sleep(1 * time.Second)
}
