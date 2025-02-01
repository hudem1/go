package main

import "fmt"

// can declare a var outside a function
var myOutsideVar = 3

// int, float, string, bool
func variables_main() {
	// compiler can infer types
	var a = "initial"
	fmt.Println("a: ", a)

	// type applies for both variables
	var b, c int = 1, 2
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)

	// type "float64"
	var fl = 2.3
	fmt.Println("fl: ", fl)

	// type "rune"
	var aChar = 'c'
	fmt.Println("aChar: ", aChar)

	var d = true
	fmt.Println("d: ", d)

	// declarations without initialization are zero-valued
	var e int
	var f int = 1 + e
	fmt.Println("f: ", f)

	var g float64 = 2
	fmt.Println("g: ", g)
	var h int = 2
	fmt.Println("h: ", h)
	// Below line gives compilation error: "invalid operation: g / h (mismatched types float64 and int)"
	// No implicit conversion for variables!!
	// var res = g/h

	// - shorthand for decl + init instead of writing: var test string = "hihi"
	// in a "normal" decl, you can omit the type, so essentially, this shorthand
	// notation removes the var keyword
	// - only usable in a function
	test := "hihi"
	fmt.Println("test: ", test)
}
