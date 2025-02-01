package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// `iptr` holds the address to a memory cell
// We tell the compiler to interpret `iptr` as a pointer to an int
func zeroptr(iptr *int) {
	*iptr = 0
}

// Pointers in go are similar to C
func pointers_main() {
	i := 1
	fmt.Println("initial:", i)

	// Does not change `i` in main because passed by value
	zeroval(i)
	fmt.Println("zeroval:", i)

	// Changes `i` in main because passed by reference
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Prints the address of `i`
	fmt.Println("pointer:", &i)
}
