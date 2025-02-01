package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func constants_main() {
	fmt.Println(s)

	// - Compiler doesn't infer explicit type right away at decl like for a var.
	// - untyped constants do not have a specific type until it is used in a context that requires one
	// 	   -> can be used in several contexts & will be given a type specific to each context
	// - here, n is an untyped integer constant
	// - It can be used with different types without explicit conversion
	// because Go allows implicit typing for constants.
	const n = 500_000_000

	// - the "literal" 3e20 is an untyped floating-point constant (float64 by default)
	// - n is an untyped integer constant.
	// - When dividing a floating-point number by n, Go implicitly promotes n to a floating-point value.
	// So, d is also an untyped floating-point constant.
	const d = 3e20 / n
	fmt.Println("d: ", d)

	// float64(d) is truncated to an int64 (not rounded).
	fmt.Println("int64 d: ", int64(d))

	// - n is an untyped integer constant,
	// so Go implicitly converts it (at compile time) to float64 when passing it to math.Sin,
	// no cast operation -> no runtime conversion
	fmt.Println("sin: ", math.Sin(n))

	// This constant flexibility in Go is powerful because it allows for type-safe
	// but flexible arithmetic without requiring manual conversions in most cases
}
