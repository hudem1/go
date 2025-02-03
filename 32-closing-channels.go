// Closing a channel indicates that no more values will be sent on it.
// This can be useful to communicate completion to the channel’s receivers.
package main

import "fmt"

func closing_channels_main() {
	// Channel to communicate work to be done from main() goroutine to a worker goroutine.
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		// infinite loop
		for {
			// This special 2-value form of receive (2nd param has just always been optional) means:
			// - j: contains the next value in the channel, otherwise the zero-value if channel is closed & empty
			// 		|--> the zero-value is not sent in the channel, but rather it is the zero-value of the type!
			// - more: is false if channel has been closed & all values in channel have been received/consumed, otherwise true
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				// Notify main goroutine all jobs are finished (27-channel-synchronization chapter)
				done <- true
				return
			}
		}
	}()

	// Send 3 jobs to worker
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	// We can close the channel even if there are remaining values to be received after -> no problem
	close(jobs)
	fmt.Println("sent all jobs")

	// Block until a value is sent on the channel (wait for worker to finish jobs)
	<-done

	// Receiving/reading from a closed channel immediately succeeds (non blocking)
	// Returned values are the same as explained above
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
