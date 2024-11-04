package main

import (
	"math/rand"
	"time"
)

func doWork(
	done <-chan interface{},
	pulseInterval time.Duration,
) (<-chan interface{}, <-chan time.Time) {
	heartbeat := make(chan interface{})
	results := make(chan time.Time)
	go func() {

		// if you want to look panic simulation comment out [defer close(ch)]
		defer close(heartbeat)
		defer close(results)

		pulse := time.Tick(pulseInterval)
		workGen := time.Tick(2 * pulseInterval)

		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
		}
		sendResult := func(r time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case results <- r:
					return
				}
			}
		}

		// if you want to look panic simulation comment out this for{} block
		for {
			select {
			case <-done:
				return
			case <-pulse:
				sendPulse()
			case r := <-workGen:
				sendResult(r)
			}
		}

		// Test Panic simuration!!
		// for i := 0; i < 2; i++ {
		// 	select {
		// 	case <-done:
		// 		return
		// 	case <-pulse:
		// 		sendPulse()
		// 	case r := <-workGen:
		// 		sendResult(r)
		// 	}
		// }
	}()
	return heartbeat, results
}

func doWork2(done <-chan interface{}) (<-chan interface{}, <-chan int) {
	heartbeatCh := make(chan interface{}, 1)
	workCh := make(chan int)

	go func() {
		defer close(heartbeatCh)
		defer close(workCh)

		for i := 0; i < 10; i++ {
			select {
			case heartbeatCh <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case workCh <- rand.Intn(10):
			}
		}
	}()
	return heartbeatCh, workCh
}

func DoWork(done <-chan interface{}, nums ...int) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{})
	intCh := make(chan int)

	go func() {
		defer close(heartbeat)
		defer close(intCh)

		time.Sleep(2 * time.Second)

		for _, n := range nums {
			select {
			case heartbeat <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case intCh <- n:
			}
		}
	}()
	return heartbeat, intCh
}

func DoWork2(
	done <-chan interface{},
	pulseInterval time.Duration,
	nums ...int,
) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	intCh := make(chan int)

	go func() {
		defer close(heartbeat)
		defer close(intCh)

		time.Sleep(2 * time.Second)

		pulse := time.Tick(pulseInterval)

		for _, n := range nums {
			select {
			case <-done:
				return
			case <-pulse:
				select {
				case heartbeat <- struct{}{}:
				default:
				}
				select {
				case <-done:
					return
				case intCh <- n:
				}

			}
		}
	}()
	return heartbeat, intCh
}
