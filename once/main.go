package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	/*
		As the name implies, sync.Once is a type that utilizes some sync primitives internally to ensure that only
		one call to Do ever calls the function passed in—even on different goroutines. This is indeed because we wrap
		the call to increment in a sync.Once Do method
	*/
	increment := func() {
		count++
	}

	var increments sync.WaitGroup
	var once sync.Once

	n := 100
	increments.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Println(count)

	// another exaple of once
	/*
		Is it surprising that the output displays 1 and not 0? This is because sync.Once only counts the number of
		times Do is called, not how many times unique functions passed into Do are called.

		I recommend that you for‐ malize this coupling by wrapping any usage of sync.Once in a small lexical block:
		either a small function, or by wrapping both in a type.
	*/
	var count int
	incrementFunc := func() { count++ }
	decrementFunc := func() { count-- }
	var once1 sync.Once
	once1.Do(incrementFunc)
	once1.Do(decrementFunc)
	fmt.Printf("Count: %d\n", count)
}
