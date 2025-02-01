// Go supports methods defined on struct types (which essentially become a parameter to the method)
//
// Methods can be defined for both pointer or value receiver types (see below methods)
//
// Use a pointer receiver to:
//   - avoid copying on method calls
//   - allow the method to mutate the receiving struct
package main

import "fmt"

type rect struct {
	width, height int
}

// Struct method takes a `*rect` as its receiver type
func (r *rect) area() int {
	return r.width * r.height
}

// Struct method takes a `rect` as its receiver type
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func methods_main() {
	r := rect{width: 10, height: 5}

	// Go automatically handles conversion between values <-> pointers receivers for method calls

	// Equivalent to (&r).area(), which passes a pointer to the method
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	// Equivalent to: (*rp).perim(), which copies the struct into a new one passed to the method
	fmt.Println("perim:", rp.perim())
}
