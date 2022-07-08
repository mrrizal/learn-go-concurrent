package example1

import (
	"fmt"
	"sync"
)

func Example() {
	// initiate the pool, in this case we will make function who print "creating new instance" reusable
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance")
			return struct{}{}
		},
	}

	// create new instance, and put into the pool
	myPool.Put(myPool.New())

	// use instance from the pool
	instance := myPool.Get()

	// when you done, just put back to the pool, so antoher process can use the instance/object
	// if not, the pool is useless
	myPool.Put(instance)

	// another process will use object/instance that availble at the pool
	myPool.Get()

	// antoher process want use the object/instance, but this time it's not availble at the pool, so in this case,
	// it will create the new one
	myPool.Get()

	// in the end, we create 2 object/instance.
}
