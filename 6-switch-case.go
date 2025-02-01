package main

import (
	"fmt"
	"time"
)

func switch_case_main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	// basic switch case
	// no indent for case statements
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// can separate statements in a case
	// (like the multiple patterns operator | in Cairo's "match" statements)
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	// like a basic if/else: with a boolean condition! Not just comparing against a single value "case 12"
	//  --> need to remove the expression following the "switch" as below
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type switch (it's an actual thing!) compares types instead of values.
	// You can use this to discover the type of an interface value.
	// An interface{} stores a value along with type information.
	whatAmI := func(i interface{}) {
		// Note that it is not an assignment in the traditional sense, instead it's a
		//  special syntax for a type switch.
		//  t is the unboxed value extracted from the interface{} cast to its actual type.
		// If you tried using i.(type) outside a switch,
		//   you'd get the error: "use of .(type) outside type switch"
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	// When a concrete type is assigned to an interface{}, Go automatically "boxes" it.
	// This is not the same as an implicit type conversion (like int to float64), but rather an interface boxing mechanism.
	// 	--> Indeed, we've seen variables don't get implicitly cast to a different type (see 2-variables.go).
	// Any type satisfies interface{} by definition.
	//
	// Those 3 calls box each arg into an interface{} type which adds a small runtime cost because:
	// - the value is stored inside a interface{} container
	// - the actual type info is tracked internally
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	// variables can be interfaces (as we saw for the argument the fct whatAmI takes)
	var here interface{}
	here = 42
	fmt.Println(here)

	// const cannot be interfaces, compiler gives error: "invalid constant type interface{}"
	// const here interface{} = 2

	var myVar = 3
	fmt.Printf("Type of myVar: %T\n", myVar)

}
