// Interfaces are named collections of method signatures (like `trait` in Cairo/Rust, `interface` in Java, etc.)
//
// In order for a struct to claim it implements an interface, it has to give an implementation
// to all methods of the interface (like usual). However unfortunately, there is no CLEAR
// indication the struct implements an interface (other than checking it implements all methods of the interface)
//
// Fortunately, the compiler will give an error if you try to use a struct (that does not implement all
// methods of an interface) where the interface is expected, but the lack of clear concise indication the struct
// implements an interface is not ideal for devX.
//
// For a struct to implement an interface, all its interface methods must take the struct by value
// (and not by pointer) as receiver. If you set the receiver as a pointer, you'll get an error saying
// the struct does not implement the interface.
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
// Must take `rect2` by value (and not by pointer). Otherwise, passing a `rect2` where a `geometry` is expected
// will result in an error like `rect2 does not implement geometry (method area2 has pointer receiver)`
func (r rect2) area2() float64 {
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
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area2())
	fmt.Println(g.perim2())
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

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
