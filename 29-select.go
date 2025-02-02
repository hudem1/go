// `select` in Go is like a switch, but for channels.
// It allows a goroutine to wait on multiple channel operations simultaneously and picks one that is ready.
package main

import (
	"fmt"
	"time"
)

func select_main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// `select` will pick the 1st case where the channel has been written to!
		// Note: it will always print first the value from `c1` (because 1 second wait only) and then from `c2`
		// The `for` loop allows to wait for both channels.
		select {
		// Note the assignment in the cases (which we didn't see with switch case)
		case msg1 := <-c1:
			fmt.Println("received from c1", msg1)
		case msg2 := <-c2:
			fmt.Println("received from c2", msg2)
			// If we were to add a default case, in both iterations the default case would be chosen
			// as both channels would not have had the time to be written to (as time sleep).
			// See `31-non-blocking-channels` chapter for more details.
			// default:
			// 	fmt.Println("No messages, moving on...")
		}
	}
}
