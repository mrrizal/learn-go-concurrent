package main

import (
	"github.com/mrrizal/learn-go-concurrent/channels/example3"
)

func main() {
	// unbuffered channel
	// example1.Example()
	// example2.Example()

	// buffered channel
	example3.Example()

	/*
		here the main different between buffered and unbuffered channels:
		-  at buffered channel, you can read the values when the channels is full or the channel is closed
			```
			sending 0
			sending 1
			sending 2
			sending 3
			sending 4
			sending 5
			received 0
			received 1
			received 2
			received 3
			received 4
			received 5
			sending 6
			sending 7
			producer done.
			received 6
			received 7
			```

	*/
}
