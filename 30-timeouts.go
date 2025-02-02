// Timeouts are useful in `select` to:
// - prevent deadlock select (if not writes are made)
// - if some (potentially external) computation takes too much time
package main

import (
	"fmt"
	"time"
)

func timeouts_main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Here the timeout case will be executed as 1 second is shorter than goroutine's 2s.
	// Note the syntax.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// Here the channel `c2` case will be executed as its 2s are shorter than timeout's 3s
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
