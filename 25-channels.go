// Channels are pipes that connect concurrent goroutines.
//
// You can send values into channels from one goroutine and
// receive those values into another goroutine.
//
// Can also simply have a variable accessible to both goroutines
// and update it in 1 goroutine, which will be reflected in the other.
// But for this method, we need:
// - some method to wait & nsure the 1st goroutine has finished updating the variable
// - a variable to be accessible by both routines.
// Also, there wouldn't be clear indicator the variable is used in both routines.
// |--> for those reasons (I think) -> better use channels.
package main

import (
	"fmt"
)

func channels_main() {
	// Create a new channel with `make(chan val-type)`. Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the `channel <-` syntax.
	// Here we send "ping" to the `messages` channel we made above, from a new goroutine.
	go func() {
		fmt.Println("before sending")
		messages <- "ping" // blocks until there is a receive (line 36)
		messages <- "pong" // blocks until there is a receive (no receiving op -> blocks undefinitely!)
		fmt.Println("after sending")
	}()

	// The `<-channel` syntax receives a value from the channel.
	// The "ping" msg is successfully passed from 1 routine to the other via the channel.
	fmt.Println("before receiving")
	msg := <-messages // blocks until there is a send (line 27)
	fmt.Println("after receiving")
	fmt.Println(msg)

	// IMPORTANT: by default, sending/receiving blocks the concerned goroutine until there is an opposite op!
	// The goroutine pauses execution at that point and waits until another goroutine is ready to complete the operation.
	// - Sending to a channel blocks the sending goroutine until another goroutine receives the value.
	// - Receiving from a channel blocks the receiving goroutine until another goroutine sends a value.
	//
	// This behaviour might create a deadlock if the main is sending/receiving but there is no opposite operation!
	// If its a "side" goroutine, it will block, true, but the main will continue. If the main blocks, the side goroutine
	// might continue and finish executing, but then we're still blocked at the main (deadlock).
}
