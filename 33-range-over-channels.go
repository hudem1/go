// We can also iterate over a channel using range, which is synonymous to receiving values from it.
//
// We have to close the channel for the range iteration to terminate. Otherwise, we'll have
// a deadlock error at execution because of the range trying to receive more values from the channel.
package main

import "fmt"

func range_over_channels_main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// `range` only returns one value at a time, the next channel value.
	// It does not return an additional optional value (whether the channel is closed & empty)
	// because the range automatically terminate when it is the case.
	for elem := range queue {
		fmt.Println(elem)
	}
}
