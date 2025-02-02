// A goroutine is a lightweight thread of execution.
// It allows to run code concurrently/asynchronously.
//
// Running goroutines might end up with different execution order (non-determinism).
//
// Note: if the main thread ends, the other ones will end too! Even if
// they haven't finished executing.
package main

import (
	"fmt"
	"time"
)

func aFct(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func go_routines_main() {
	// Run synchronously, like usual
	aFct("direct")

	// Just need to add keyword `go` to invoke the fct in a goroutine
	go aFct("goroutine")

	// Can also start a goroutine for an anonymous fct
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// The 2 above calls will run asynchronously in separate goroutines.

	// Make the main thread wait 1 sec to wait for other threads
	// to finish their execution. Obviously, this method is not ideal,
	// and we can instead use `WaitGroup` to wait for them to finish.
	time.Sleep(time.Second)
	fmt.Println("done")
}
