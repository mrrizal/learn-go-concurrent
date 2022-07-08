package main

import (
	"fmt"

	"github.com/mrrizal/learn-go-concurrent/pool/example1"
	"github.com/mrrizal/learn-go-concurrent/pool/example2"
)

func main() {
	/*
		the idea of pool is to make object or instance reusable.

		the main benefit to use pool is reduce memory usage make your program faster.

		it's becuase you no need to create much object or instance, if the object or instance is available at the pool,
		just reuse it, if not then create the new one.
	*/

	example1.Example()
	fmt.Println("=========EXAMPLE 2=========")
	example2.Example()
}
