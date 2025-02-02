// Interfaces are named collections of method signatures (like `trait` in Cairo/Rust, `interface` in Java, etc.)
//
// In order for a struct to claim it implements an interface, it has to give an implementation
// to all methods of the interface (like usual). However unfortunately, there is no CLEAR
// indication the struct implements an interface (other than checking it implements all methods of the interface)
//
// Fortunately, the compiler will give an error if you try to use a struct (that does not implement all
// methods of an interface) where the interface is expected, but the lack of clear concise indication the struct
// implements an interface is not ideal for devX.
package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area2() float64
	perim2() float64
}

type rect2 struct {
	width, height float64
}

type circle struct {
	radius float64
}

// `area2` is a method on rect2
// Fyi, passing the receiver by value or pointer has no incidence on the interface implementation claim --> both work
func (r *rect2) area2() float64 {
	return r.width * r.height
}

// `perim2` is a method on rect2
func (r rect2) perim2() float64 {
	return 2*r.width + 2*r.height
}

// `area2` is a method on circle
func (c circle) area2() float64 {
	return math.Pi * c.radius * c.radius
}

// `perim2` is a method on circle
func (c circle) perim2() float64 {
	return 2 * math.Pi * c.radius
}

// Can pass anything that implements the `geometry` interface
//
// Fyi, passing a pointer to an interface would not really make sense, because interfaces
// are already abstractions, they don't have a concrete representation in memory like a struct at compile-time
// |--> at compile time: its size is not known, neither its memory layout (alignmnent, padding, etc.)
// Only at runtime do they have a memory representation: a pair `(concrete type, concrete type's values)`
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area2())
	fmt.Println(g.perim2())

	// Just for information purposes
	fmt.Printf("runtime geometry interface g's pair: %T, %v\n", g, g)
	if value, ok := g.(*rect2); ok {
		// address/pointer (stored in interface) to the original `rect2`
		fmt.Printf("Value inside interface: %p\n", value)
	}
}

func detectCircle(g geometry) {
	// Sometimes it's useful to know the runtime type of an interface value
	//
	// One option is using a type assertion as shown here (only checks against 1 type)
	// Another is a type switch (get the actual type and perform several type assertions -> switch-case)
	//
	// `g.(type)` can be used only in type-switch, whereas `g.(aType)` can be used anywhere
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func interfaces_main() {
	r := rect2{width: 3, height: 4}
	c := circle{radius: 5}

	// Need to pass a pointer to `r` because method `area2` for `r` takes its receiver by pointer.
	//
	// Even if Go in the general case handles automatic conversion between value <-> pointer when calling
	// a method on a type, here as the method is being called through the intermediary of an interface,
	// Go can still automatically deference the pointer to `rect2` when calling `perim2`, but if we were to pass
	// `r` by value to `measure`, Go would not be able to convert `rect2` to a pointer when calling `area2`.
	// The reason is: when assigning the concrete type to the interface variable, a copy of the concrete type's
	// value is copied into the interface tuple! Therefore, once copied, the interface tuple no longer has
	// acccess to the original struct and sees only the copy! So, the compiler cannot access the original struct's address
	// and therefore cannot auto-promote the interface value to a pointer (which would only reference the interface's
	// struct and not the original one) !!
	measure(&r)
	measure(c)

	detectCircle(&r)
	detectCircle(c)
}
