// By default channels are unbuffered, meaning that they will only accept sends (chan <-)
// if there is a corresponding receive (<- chan) ready to receive the sent value (as seen in previous chapter).
//
// Buffered channels accept up to a certain number of values without a corresponding receiver for those values.
//
// However, if there is no sent value but a receiving operation, it will block! Buffer is only for sends, not receives!!
package main

import "fmt"

func buffered_channels_main() {
	messages := make(chan string, 2)

	// We have 2 sends, so, its fine, execution won't block
	// If we added a 3rd send, exec would block and we'd get a deadlock.
	messages <- "buffered"
	messages <- "channel"

	// If the main had only this receiving op with no sending op before, it would block!
	// Buffer is only for sends (as explained at the beginning of file).
	// <-messages

	// Don't even need those 2 lines for the program to correctly execute.
	// The only thing is values passed to the channel would not be received, but
	// that's not necessarily an issue (at least not compilation/exec wise!)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
