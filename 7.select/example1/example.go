package example1

import (
	"fmt"
	"time"
)

func Example() {
	/*
		this example will block select statement for 5 seconds
		it's because the select statement will only executed when there is a data sending to that channel
		and the select statement read it, or the channel has been closed.

		in this case, there is no data will be sent into the channel, so the select statement will be executed
		because of channel is closed instead.
	*/
	start := time.Now()
	c := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("blocking on read")
	select {
	case <-c:
		fmt.Printf("unblocked %v later.\n", time.Since(start))
	}

}
