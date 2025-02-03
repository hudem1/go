// We can use channels to synchronize execution across goroutines.
// Here’s an example of using a blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
package main

import (
	"fmt"
	"time"
)

// Note how to pass the channel as function param `chan chan-type`
func workerCh27(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Notify end of exec
	done <- true
}

func channel_synchronization_main() {
	done := make(chan bool, 1)
	go workerCh27(done)

	// Block until we receive a sending op (end-of-exec notification)
	<-done
}
