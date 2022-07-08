package example2

import (
	"fmt"
	"sync"
)

func Example() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup

	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created\n", numCalcsCreated)

	fmt.Println("not put instance back to pool")
	numCalcsCreated = 0
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			_ = calcPool.Get().(*[]byte)
			// let's not put the instance to pool
			// defer calcPool.Put(mem)
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created\n", numCalcsCreated)
}
