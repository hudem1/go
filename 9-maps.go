package main

import (
	"fmt"
	"maps"
)

func maps_main() {
	// Creates an empty map --> map[key-type]value-type
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	// Prints all entries
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If key does not exist, return value-type's zero-valued (here empty string "")
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// Number of entries
	fmt.Println("len:", len(m))

	// Can delete an entry
	delete(m, "k2")
	fmt.Println("map:", m)

	// Can delete all map entries
	clear(m)
	fmt.Println("map:", m)

	// Accessing a map key returns 2 things: value, was_key_present (a bool)
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Decl + init in 1 line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	// Map package contains lots of utility functions
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
