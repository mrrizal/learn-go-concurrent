package main

import "github.com/mrrizal/learn-go-concurrent-cond/example2"

func main() {
	/*
		https://developpaper.com/detailed-explanation-of-sync-cond-in-go-language/
		Cond in golang’s sync package implements a condition variable that can use multiple readers to wait
		for public resources (using broadcast).

		Each cond is associated with a lock. When modifying a condition or calling the wait method, it must be locked
		to protect the condition. It is similar to wait and notifyAll in Java.

		The sync.cond condition variable is used to coordinate those goroutines that want to share resources.
		When the state of shared resources changes, it can be used to notify goroutines blocked by mutex.
	*/

	// this one using signal
	// example1.Example()

	// this using broadcast
	example2.Example()
}
