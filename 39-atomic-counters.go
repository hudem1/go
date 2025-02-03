// The primary mechanism for managing state when using Goroutines is communication over channels.
// But there are a few other options for managing state.
//
// Here we’ll look at using the `sync/atomic` package for atomic counters accessed by multiple goroutines.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomic_counters_main() {
	// Atomic integer.
	var ops atomic.Uint64

	// To wait for all routines to finish executing.
	var wg sync.WaitGroup

	// Create 50 routines.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// Each routine increments the counter by 1 a 1000 times.
		go func() {
			for c := 0; c < 1000; c++ {
				// Use method `Add` to atomically increment the counter (not `++`, we can't anyway on `atomic.Uint64`)
				ops.Add(1)
			}

			// Note here we call `Done` without using `defer` (there is no worry of the fct exiting early).
			wg.Done()
		}()
	}

	wg.Wait()

	// Here all goroutines are done writing to `ops`.
	// But atomically reading a value using `Load` is safe even while other goroutines are (atomically) updating it.
	//
	// We indeed get a result of 50_000 as expected. Had we used non-atomic integer, we'd likely get a different number
	// changing between runs because goroutines would interfere with each other. Moreover, we'd get data race failures
	// when running with the `-race` flag.
	fmt.Println("ops:", ops.Load())
}
