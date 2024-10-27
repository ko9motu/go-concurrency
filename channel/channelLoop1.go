package channel

import (
	"fmt"
)

func ChannelLoop1() {
	intCh := make(chan int)
	go func() {
		defer close(intCh)

		for i := 1; i <= 5; i++ {
			intCh <- i
		}
	}()

	for integer := range intCh {
		fmt.Printf("%v ", integer)
	}
}
