package channel

import (
	"fmt"
	"sync"
)

func ChannelLoop2() {
	begin := make(chan interface{})

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unbroking goroutines...")
	close(begin)
	wg.Wait()
}
