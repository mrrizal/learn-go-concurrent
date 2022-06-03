package main

import (
	"fmt"
	"sync"
	"time"
)

func example1() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping")
		time.Sleep(1 * time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping")
		time.Sleep(2 * time.Second)
	}()

	wg.Wait()
	fmt.Println("All goroutines completes.")
}

// commonly used
func example2() {
	var wg sync.WaitGroup

	myGoroutine := func(id int) {
		defer wg.Done()
		fmt.Printf("hello from goroutine %d\n", id)
	}

	n := 5
	wg.Add(n)
	for i := 0; i <= n; i++ {
		go myGoroutine(i)
	}
	wg.Wait()
	fmt.Println("All goroutines completes.")
}

func main() {
	example1()
	example2()
}
