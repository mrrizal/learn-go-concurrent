package example4

import "fmt"

func UniDirectionalChannel() {
	// you can create directional channel with this way
	var dataStream = make(chan interface{})
	var receivedStream <-chan interface{}
	var sendStream chan<- interface{}

	receivedStream = dataStream
	sendStream = dataStream

	go func() {
		defer close(sendStream)
		for i := 0; i < 10; i++ {
			dataStream <- i
		}
	}()

	for data := range receivedStream {
		fmt.Printf("received :%d\n", data)
	}

}
