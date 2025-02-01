// Go doesn't have an enum type as a distinct lang feature, but enums
// can be implemented easily.
package main

import "fmt"

type ServerState int

// Creates possible values for ServerState defined as constants.
// Keyword `iota` generates successive constant values automatically (here, 0, 1, 2, 3)
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// Issue: does not enforce value for all ServerState keys (can omit a ServerState key) -> error prone
// Issue: can add any int key --> error prone
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// - Below method allows to claim ServerState implements `fmt.Stringer` interface,
// allowing values of ServerState to be printed out or converted to strings
// - Note that below method is implemented on a type that is not a struct!
func (ss ServerState) String() string {
	return stateName[ss]
}

func enums_main() {
	// Cannot pass an int as the argument to `transition` --> compiler will give an error
	//  |--> Provides some degree of compile-time type safety for enums
	//  |--> Passing a variable of type int raises a compiler error
	//       but passing an untyped int (constant) `2` by itself does NOT.
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		// s will be empty, will not print anything (because no ServerState matches its value --> i tried with constant > 3)
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
