package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_doWork(t *testing.T) {
	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() { close(done) })

	const timeout = 2 * time.Second
	heartbeat, results := doWork(done, timeout/2)

	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if ok == false {
				return
			}
			fmt.Printf("results %v\n", r.Second())
		case <-time.After(timeout):
			fmt.Println("worker goroutine is not healty")
			return
		}
	}
}

func Test_doWork2(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	heartbeat, results := doWork2(done)

	for {
		select {
		case <-done:
			return
		case _, ok := <-heartbeat:
			if ok {
				fmt.Println("pulse")
			} else {
				return
			}
		case r, ok := <-results:
			if ok {
				fmt.Printf("results %v\n", r)
			} else {
				return
			}
		}
	}

}

func Test_Bad_DoWork_GenerateAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	_, results := DoWork(done, intSlice...)

	for i, expected := range intSlice {
		select {
		case r := <-results:
			if r != expected {
				t.Errorf(
					"index %v: expected %v, but received %v", i, expected, r,
				)
			}
		case <-time.After(1 * time.Second):
			t.Fatal("test timed out")
		}
	}
}

func Test_Good_DoWork_GenerateAllNumbers(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	hearbeat, results := DoWork(done, intSlice...)

	<-hearbeat

	i := 0
	for r := range results {
		if expected := intSlice[i]; r != expected {
			t.Errorf("index %v: expected %v, but recieved %v,", i, expected, r)
		}
		i++
	}
}

func Test_DoWork2(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}

	const timeout = 2 * time.Second

	heartbeat, results := DoWork2(done, timeout/2, intSlice...)

	<-heartbeat

	i := 0

	for {
		select {
		case r, ok := <-results:
			if ok == false {
				return
			} else if expected := intSlice[i]; r != expected {
				t.Errorf(
					"index %v: expected: %v, but received %v",
					i,
					expected,
					r,
				)
			}
			i++
		case <-heartbeat:
		case <-time.After(timeout):
			t.Fatal("test timed out")
		}
	}
}
