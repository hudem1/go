// Go’s built-in `ticker` is used when you want to do something repeatedly at regular intervals.
package main

import (
	"fmt"
	"time"
)

func tickers_main() {
	// Similar mechanism to `timers` but `tickers` trigger an action every x time (here 500ms).
	//
	// Under the hood, `ticker` creates a buffered channel with a capacity of 1. Just as we've seen with buffered channels,
	// once full and trying to send 1 more, it will not send it. But before, we've seen that it blocks execution at the extra send
	// and once unblocked, it will send the extra send it was blocking at. Here, it doesn't really block at the extra tick,
	// instead the extra tick not only would not get sent (normal), but also get dropped (new) !
	// https://stackoverflow.com/questions/71191067/golang-time-ticker-triggers-twice-after-blocking-a-while
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Pause the root goroutine enough time for 3 ticks
	time.Sleep(1600 * time.Millisecond)

	// Stop the `ticker`, which stops firing ticks.
	//
	// Note: stopping the ticker does not close the channel. Otherwise, as we've seen before,
	// the receiver would read (not really "receive" anymore as it's not sent) zero-value `time.Time{}`
	// (after all previous tick values have been received). And, we don't want the receiver to read values
	// that were not ticks and misinterpret for ticks.
	ticker.Stop()

	done <- true
	fmt.Println("Ticker stopped")
}
