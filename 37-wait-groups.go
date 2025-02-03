// `WaitGroup` is used to wait for multiple goroutines to finish.
//
// It provides 3 main methods:
// - Add(int): increases the counter by the given number (representing the number of goroutines to wait for).
// - Done(): decreases the counter by one (should be called in the goroutine when it's done).
// - Wait(): blocks execution until the counter becomes zero.
//
// Note that this approach has no straightforward way to propagate errors from workers.
// For more advanced use cases, consider using the `errgroup` package.
package main

import (
	"fmt"
	"sync"
	"time"
)

func workerCh37(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func wait_groups_main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// Increment the counter by 1.
		wg.Add(1)

		go func() {
			// Defer the call (to Done()) to the end of this function
			// meaning the counter gets decremented when the goroutine terminates.
			//
			// We could also call `Done` manually (without `defer`) at the end of the `workerCh37` or this fct,
			// but `defer`:
			// - ensures the call to `Done` happpens even if the function exits due to an error or an early return,
			// - avoids for the worker itself (the fct being exec) to be aware of the concurrency primitives involved in its exec.
			defer wg.Done()

			// If we were to pass the `waitGroup` to the function (to maybe call `Done` inside or smth),
			// always pass it by pointer! Otherwise, each goroutine gets a copy, and the main `WaitGroup`
			// won't track them and who's done.
			workerCh37(i)
		}()
	}

	// Block execution until the underlying counter to `WaitGroup` reaches 0,
	// meaning all goroutines finished executing.
	//
	// We shouldn't reuse the same `WaitGroup` after the call to `Wait`, without reinitializing it.
	wg.Wait()
}
