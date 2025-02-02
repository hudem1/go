// In Go, errors are communicated via an explicit separate return value.
//
// By convention, errors are the last return value and have type `error`, a built-in interface.
package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		// `errors.New` constructs a basic `error` value with the given error msg
		return -1, errors.New("can't work with 42")
	}

	// `nil` means there was no error
	return arg + 3, nil
}

// `Sentinel` errors are predeclared variables used to signify specific error values
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		// We can wrap errors with higher-level errors to add context.
		// Creates a logical chain (A wraps B that wraps C, etc.), which can be
		// queried with functions like `errors.Is` and `errors.As`
		//
		// `%w` is used to ensure the passed error `ErrPower` is preserved and can be unwrapped
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

// Also have the possibility to create custom errors (convention wants suffix `Error`)
// by implementing the `Error()` method on them --> `error` interface
type argError struct {
	arg     int
	message string
}

// Thanks to implementing this method (`error` interface type)
// `argError` type is now considered an `error` type as well
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// Can now return our customError where an `error` is expected!
func someFct(val int) error {
	if val >= 0 {
		return nil
	}

	// Even though Go can usually automatically convert value <-> pointer for
	// receivers for method calls, here need to explicitly pass a pointer to type
	// (as we passed a pointer receiver in `Error` method definition).
	// See `18-interface` chapter for more details.
	return &argError{arg: 0, message: "Negative number!"}
}

func errors_main() {
	for _, i := range []int{7, 42} {
		// Common to use an inline error check in `if`
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			// `errors.Is` checks that a given error (or any error in its chain)
			// matches a specific error value. This is especially useful with wrapped
			// or nested errors, allowing to identify specific error types or sentinels
			// in a chain of errors.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}

	err := someFct(-1)
	var ae *argError
	// `errors.As` is a more advanced version of `errors.Is`. Not only does it do the
	// same check as `errors.Is` (verify that it matches a specific error TYPE here!),
	// but also, if that check passes, sets the 2nd parameter the value of that type.
	//
	// In other words, if the `As` call succeeds (return `true`), var `ae`
	// will point to the concrete (potentially nested) error!
	if errors.As(err, &ae) {
		fmt.Println(ae.arg, "-", ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}

	// Below code compiles but fails (panics) at exec because `As` expects a pointer to a
	// variable of the type you're trying to match against, not the basic `error` type!
	// Compilation passes because it seems like type checking passes (as `error` is of
	// any interface type).
	//
	// var otherErr = fmt.Errorf("Composed error: %w", ae)
	// if errors.As(err, otherErr) {
	// 	// do smth
	// }
}
