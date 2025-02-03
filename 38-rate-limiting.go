// Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service.
// Go elegantly supports rate limiting with goroutines, channels, and tickers.
package main

import (
	"fmt"
	"time"
)

func rate_limiting_main() {
	// 1. Basic limiting

	// Create 5 requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	// Treat every request at a rate of 200ms.
	//
	// As a reminder: when iterating over `requests`, the root routine receives
	// requests immediately from the channel as they were already sent before.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 2. Burst limiting

	// The idea is to allow short bursts of requests in our rate limiting scheme
	// while preserving the overall rate limit.
	//
	// This `burstyLimiter` channel will allow bursts of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		// Infinite loop with a value (`time.Time`) at every 200ms.
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Create some requests for this 2nd example.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	// As a reminder: close the channel to ensure termination for the `range` iteration.
	close(burstyRequests)

	// Iterate over all requests, and the 3 first ones will be received/served immediately
	// while the next ones will be served at a rate of 200ms.
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
