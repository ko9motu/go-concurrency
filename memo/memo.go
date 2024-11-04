package memo

import "fmt"

func Gen(ints ...int) <-chan int {
	intCh := make(chan int)

	go func() {
		defer close(intCh)

		for _, i := range ints {
			intCh <- i
		}
	}()
	return intCh
}

func main() {

	intCh := Gen(1, 2, 3, 4)
	for i := 0; i <= 3; i++ {
		fmt.Println(<-intCh)
	}

}
