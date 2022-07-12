package example1

import "fmt"

func UnBufferedChannel() {
	// unbuffered channel
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		defer fmt.Println("producer done.")
		for i := 0; i < 5; i++ {
			fmt.Printf("sending %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("received %d\n", integer)
	}
}
