package example1

import (
	"fmt"
	"sync"
	"time"
)

func Example() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		// it's used for notifying goroutine blocked on a Wait call that the condition has been triggered.
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		if len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
	fmt.Printf("remaining data %d\n", len(queue))
}
