package channel

import (
	"bytes"
	"fmt"
	"os"
)

func Buffer() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intCh := make(chan int, 4)
	go func() {
		defer close(intCh)
		defer fmt.Println(&stdoutBuff, "Produecer done")

		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "sending: %v.\n", i)
			intCh <- i
		}
	}()

	for integer := range intCh {
		fmt.Fprintf(&stdoutBuff, "Received: %v.\n", integer)
	}
}
