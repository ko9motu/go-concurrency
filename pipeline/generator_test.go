package pipeline

import "testing"

func BenchmarkGeneric(b *testing.B) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for range ToString(done, Take(done, Repeat(done, "a"), b.N)) {
	}
}

func BenchmarkTyped(b *testing.B) {
	repeat := func(done <-chan interface{}, values ...string) <-chan string {
		valueCh := make(chan string)
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

	take := func(done <-chan interface{}, valueCh <-chan string, num int) <-chan string {
		takeCh := make(chan string)
		go func() {
			defer close(takeCh)
			for i := num; i > 0 || i == -1; {
				if i != -1 {
					i--
				}
				select {
				case <-done:
					return
				case takeCh <- <-valueCh:
				}
			}
		}()
		return takeCh
	}
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for range take(done, repeat(done, "a"), b.N) {
	}
}
