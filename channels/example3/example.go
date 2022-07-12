package example3

import "fmt"

func BufferedChannel() {
	intStream := make(chan int, 4)
	fmt.Println(cap(intStream))

	go func() {
		defer close(intStream)
		defer fmt.Println("producer done.")
		for i := 0; i < 7; i++ {
			fmt.Printf("sending %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("received %d\n", integer)
	}
}
