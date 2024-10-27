package pipeline

import "fmt"

// b or B is mean of Batch
func b_multiply(values []int, multiplier int) []int {
	multipliedValues := make([]int, len(values))
	for i, v := range values {
		multipliedValues[i] = v * multiplier
	}
	return multipliedValues
}

func b_add(values []int, additive int) []int {
	addValues := make([]int, len(values))
	for i, v := range values {
		addValues[i] = v + additive
	}
	return addValues
}

func B_pipeline() {
	ints := []int{1, 2, 3, 4}
	for _, v := range b_add(b_multiply(ints, 2), 1) {
		fmt.Println(v)
	}
}

// s or S is mean of Stream
func s_multiply(value int, multiplier int) int {
	return value * multiplier
}

func s_add(value int, additive int) int {
	return value + additive
}

func S_pipeline() {
	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Println(s_multiply(s_add(s_multiply(v, 2), 1), 2))
	}
}

// g or G is mean of Good pattern when I use "Go"
func g_generator(done <-chan interface{}, integers ...int) <-chan int {
	intCh := make(chan int, len(integers))
	go func() {
		defer close(intCh)
		for _, i := range integers {
			select {
			case <-done:
				return
			case intCh <- i:
			}
		}
	}()
	return intCh
}

func g_multiply(done <-chan interface{}, intCh <-chan int, multiplier int) <-chan int {
	multipliedCh := make(chan int)
	go func() {
		defer close(multipliedCh)
		for i := range intCh {
			select {
			case <-done:
				return
			case multipliedCh <- i * multiplier:
			}
		}
	}()
	return multipliedCh
}

func g_add(done <-chan interface{}, intCh <-chan int, additive int) <-chan int {
	addedCh := make(chan int)
	go func() {
		defer close(addedCh)
		for i := range intCh {
			select {
			case <-done:
				return
			case addedCh <- i * additive:
			}
		}
	}()
	return addedCh
}

func G_pipline() {
	done := make(chan interface{})
	defer close(done)

	intCh := g_generator(done, 1, 2, 3, 4)
	pipeline := g_multiply(done, g_add(done, g_multiply(done, intCh, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}
