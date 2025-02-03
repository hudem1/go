// Go’s built-in `timer` is used when you want to do something once in the future.
package main

import (
	"fmt"
	"time"
)

func timer_main() {
	// Timers represent a single event in the future. Here, it triggers an action after 2s.
	// It returns an object that provides a channel that will be notified after that time.
	timer1 := time.NewTimer(2 * time.Second)

	// As always, receiving from a channel `.C` is blocking.
	// Wait until a message is sent to the channel (time is elapsed).
	//
	// Could also have used a `time.Sleep` to wait but a timer is more powerful
	// as you can trigger the timer somewhere and be notified for the elapsed time somewhere else.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// Can cancel the timer before it fires!
	//
	// The above routine was waiting for the timer to trigger after 1s, but as we stopped it,
	// the routine will get blocked on trying to receive and not finish its execution.
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// We add a 2s wait just to show that the side routine is indeed blocked
	// and never reaches "Timer 2 fired" code.
	time.Sleep(2 * time.Second)
}
