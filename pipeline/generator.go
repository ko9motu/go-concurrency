package pipeline

import (
	"fmt"
	"math/rand"
)

func Repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueCh := make(chan interface{})
	go func() {
		defer close(valueCh)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueCh <- v:
				}
			}
		}
	}()
	return valueCh

}

func Take(done <-chan interface{}, valueCh <-chan interface{}, num int) <-chan interface{} {
	takeCh := make(chan interface{})
	go func() {
		defer close(takeCh)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeCh <- <-valueCh:
			}
		}
	}()
	return takeCh
}

func RepeatAndTakePipeline() {
	done := make(chan interface{})
	defer close(done)

	for num := range Take(done, Repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}
}

func ExRepeat(done <-chan interface{}, function func() interface{}) <-chan interface{} {
	valueCh := make(chan interface{})
	go func() {
		defer close(valueCh)
		for {
			select {
			case <-done:
				return
			case valueCh <- function():
			}
		}
	}()
	return valueCh
}

func UseCaseExRepeat() {
	done := make(chan interface{})
	defer close(done)

	rand := func() interface{} { return rand.Int() }

	for num := range Take(done, ExRepeat(done, rand), 10) {
		fmt.Println(num)
	}
}

func ToString(done <-chan interface{}, valueCh <-chan interface{}) <-chan string {
	stringCh := make(chan string)
	go func() {
		defer close(stringCh)
		for v := range valueCh {
			select {
			case <-done:
				return
			case stringCh <- v.(string):
			}
		}
	}()
	return stringCh
}

func UseCaseToString() {
	done := make(chan interface{})
	defer close(done)
	var message string
	for token := range ToString(done, Take(done, Repeat(done, "I", "am."), 5)) {
		message += token
	}
	fmt.Printf("message: %s...", message)
}
