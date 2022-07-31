package example4

import (
	"fmt"
	"time"
)

func Example() {
	// default case, normally use inside for loop to make a progress
	start := time.Now()
	var c chan interface{}
	select {
	case <-c:
		fmt.Println("receive data from channel")
	default:
		fmt.Println("done")
	}

	fmt.Printf("exited after %v\n", time.Since(start))

	channel := make(chan int)
	counter := 0

	go func() {
		time.Sleep(5 * time.Second)
		close(channel)
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel <- 100
	}()

	// please take a note, when we close the channel, the channel will return empty data depens on what channel type is
	// loop:
	// 	for {
	// 		select {
	// 		case data := <-channel:
	// 			fmt.Println(data)
	// 			break loop
	// 		default:
	// 		}

	// 		time.Sleep(1 * time.Second)
	// 		counter++
	// 	}

	// another way doing for loop
	for {
		select {
		case data := <-channel:
			fmt.Println(data)
			if data == 0 {
				fmt.Printf("counter: %d\n", counter)
				return
			}
		default:
		}
		time.Sleep(1 * time.Second)
		counter++
	}
}
