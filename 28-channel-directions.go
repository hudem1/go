// When using channels as function parameters, you can specify if a channel is meant to only send or receive values.
//
// This specificity increases the type-safety of the program. A compilation error would arise were we
// to try to receive in a fct where the channel is passed as send-only, or vice versa.
package main

import "fmt"

// This `ping` function only accepts a channel for sending values.
// It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	// No blocking whatsoever (as it's a send + buffered channel + no other pings send)
	pings <- msg
}

// The `pong` function accepts one channel for receives (`pings`) and a second for sends (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
	// 2. Blocks until a ping send
	msg := <-pings
	pongs <- msg
}

func channel_directions_main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	go ping(pings, "passed message")
	go pong(pings, pongs)

	// 1. Blocks until a pong send
	fmt.Println(<-pongs)
}
