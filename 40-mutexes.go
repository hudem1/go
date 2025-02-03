// In the previous example we saw how to manage simple counter state using atomic operations.
//
// For more complex state we can use a mutex to safely access data across multiple goroutines.
package main

import (
	"fmt"
	"sync"
)

// Container holds a map of counters; since we want to update it concurrently
// from multiple goroutines, we add a Mutex to synchronize access.
//
// Note that mutexes must not be copied, so if this struct is passed around, it should be done by pointer!
// Otherwise, routines would update the same map (because for the map field, copying the `Container` struct
// would only copy the map header that holds a pointer to the data) but with each routine using a different mutex.
// The resulting map would not necessarily be what we expect (nor determinist for each run).
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	// Could also call `Unlock` manually (without `defer`) after updating the `counters` map.
	// `defer` is used just for safety in case the function is updated someday with multiple possible exits for ex.
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	// Note that the zero value of a mutex is usable as-is, so no initialization is required here.
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	// As a reminder, as the variable is a struct with counters, their zero-value initialization is
	// completely fine and works well (without needing a manual initialization).
	var wg sync.WaitGroup

	// The worker goroutines will call.
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	// All 3 goroutines will update the same `Container` object, including 2 of them updating the same counter.
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
