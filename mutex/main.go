package main

import (
	"fmt"
	"time"

	"github.com/mrrizal/learn-go-concurrent/mutex/example1"
	"github.com/mrrizal/learn-go-concurrent/mutex/example2"
	"github.com/mrrizal/learn-go-concurrent/mutex/example3"
)

func main() {
	start := time.Now()
	// not using mutex, result will be inconsistent
	for i := 0; i < 10; i++ {
		example1.Example()
	}
	fmt.Printf("elapsed: %s\n", time.Since(start))

	fmt.Println("=======================")

	start = time.Now()
	// use mutex
	for i := 0; i < 10; i++ {
		example2.Example()
	}
	fmt.Printf("elapsed: %s\n", time.Since(start))

	fmt.Println("=======================")

	start = time.Now()
	// use rwmutex
	/*
		https://dev.to/cristicurteanu/mutexes-and-rwmutex-in-golang-4ij#:~:text=The%20most%20straightforward%20approach%20would,become%20ineffective%20to%20use%20RWMutex%20.
		why we need RWMutex?
		That's fair enough. But the problem is, that when using Mutex the value from the memory will be locked until the
		Unlock method will be invoked. This is also valable for the reading phase. In order to make reading accessible
		for multiple threads, the Mutex can be replaced with RWMutex, and for reading it will be used RLock and RUnlock methods.
	*/

	// https://stackoverflow.com/questions/48861029/what-is-the-benefit-of-using-rwmutex-instead-of-mutex
	/*
		In other words, readers don't have to wait for each other. They only have to wait for writers holding the lock.
		A sync.RWMutex is thus preferable for data that is mostly read, and the resource that is saved compared to a sync.Mutex is time.
	*/
	for i := 0; i < 10; i++ {
		example3.Example()
	}
	fmt.Printf("elapsed: %s\n", time.Since(start))
}
