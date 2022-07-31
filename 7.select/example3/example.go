package example3

import (
	"fmt"
	"time"
)

func Example() {
	/*
		this select statement will never unblock because:
		1. the channel is empty/nil
		2. channel never closed

		so in this case, the select statement will be exited after 2 second
	*/
	var c chan interface{}
	select {
	case <-c:
		fmt.Println("read data from channel")
	case <-time.After(2 * time.Second):
		fmt.Println("timed out")
	}
}
