package example2

import (
	"log"
	"sync"
	"time"
)

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	log.Println(name, "start reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "start writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")

	// when broadcat called, it will unlock all goroutine
	c.Broadcast()
}

func Example() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader 1", cond)
	go read("reader 2", cond)
	go read("reader 3", cond)

	write("writer 1", cond)
	time.Sleep(time.Second * 3)
}
