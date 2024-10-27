package concurrencypattern

import (
	"fmt"
	"time"
)

func GoroutineLeak() {
	dowork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("dowork exited")
			defer close(completed)

			for s := range strings {
				// do something...
				fmt.Println(s)
			}
		}()
		return completed
	}
	dowork(nil)
	// do something...
	fmt.Println("Done.")
}

func AvoidGoroutineLeak() {
	dowork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("dowork exited.")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					// do something
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := dowork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancel dowork goruitne!")
		close(done)
	}()

	<-terminated
	fmt.Println("Done")
}
