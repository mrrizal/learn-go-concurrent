package main

import (
	"fmt"

	"github.com/mrrizal/learn-go-concurrent/mutex/example1"
	"github.com/mrrizal/learn-go-concurrent/mutex/example2"
)

func main() {
	for i := 0; i < 10; i++ {
		example1.Example()
	}

	fmt.Println("=======================")

	for i := 0; i < 10; i++ {
		example2.Example()
	}
}
