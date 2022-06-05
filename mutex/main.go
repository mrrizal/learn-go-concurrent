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
	for i := 0; i < 10; i++ {
		example3.Example()
	}
	fmt.Printf("elapsed: %s\n", time.Since(start))
}
