package pipeline

import "fmt"

func Bridge(
	done <-chan interface{},
	chStream <-chan <-chan interface{},
) <-chan interface{} {
	valCh := make(chan interface{})
	go func() {
		defer close(valCh)
		for {
			var ch <-chan interface{}
			select {
			case maybeCh, ok := <-chStream:
				if ok == false {
					return
				}
				ch = maybeCh
			case <-done:
				return
			}
			for val := range OrDone(done, ch) {
				select {
				case valCh <- val:
				case <-done:
				}
			}
		}
	}()
	return valCh
}

func GenVals() <-chan <-chan interface{} {
	chStream := make(chan (<-chan interface{}))
	go func() {
		defer close(chStream)
		for i := 0; i < 10; i++ {
			ch := make(chan interface{})
			ch <- i
			close(ch)
			chStream <- ch
		}
	}()
	return chStream
}

func CheckBridge() {
	for v := range Bridge(nil, GenVals()) {
		fmt.Printf("%v: ", v)
	}
}
