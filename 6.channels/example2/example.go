package example2

import (
	"fmt"
	"sync"
)

func Example() {
	// unblocking multiple goroutines
	begin := make(chan interface{})
	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking all goroutines")
	close(begin)
	wg.Wait()
}
