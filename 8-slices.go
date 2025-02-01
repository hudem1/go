package main

import (
	"fmt"
	"slices"
)

func slices_main() {
	// Slices are typed only by the elements they contain (not the length) -> in comparison to arrays!
	// An uninitialized slice equals to nil (because Go's zero-value for slices is nil) and has length 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// `make` allocates & initializes an object of type slice (therefore not nil).
	// Initialized slice with 3 zero-valued strings!
	// By default, capacity == length. But can add capacity as a 3rd param if we know the
	// slice is gonna grow in size.
	s = make([]string, 3)
	fmt.Println("emp:", s, s != nil, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	// Slices can be copied --> Shallow copy: the elements are copied by value (safe for primitive types).
	// If the elements are references (slices, maps, pointers, strings, etc?), then only the ref are copied!
	//
	// Even though strings are shallow copied, modifying c or s will not alter another!
	// The reason is strings are immutable, they cannot be modified in place! They implement copy-on-write,
	// meaning a new string is allocated only when it's written to, like the example below:
	// Ex: c[0] = "mystr" --> a new string would be created!
	// But a mere: myStr2 := myStr1 does not create a new string, but copies only the header (pointer-to-data + len + cap)
	//
	// When slices are copied, a new slice structure is allocated holding the same values as the src's:
	// - the 3 fields: pointer-to-data, length and capacity are copied
	// - in addition, the pointer-to-data is not actually simply copied to point to the src's data, but new slice data
	//   is allocated to hold the same references (hence shallow copy) as the src's data.
	// In summary, the underlying data (here strings) is shallow copied! But a new slice structure (small size) is created.
	//
	// 			Scenario									Copy Type
	// copy([]string, []string)					Shallow Copy (only references copied)
	// []byte(s) (string → byte slice)			Deep Copy (new byte slice allocated)
	// string([]byte) (byte slice → string)		Zero-Cost Conversion (reuses data, just adds a string header)
	copy(c, s)
	fmt.Println("cpy:", c)

	// Can use the "slice" operator on slices [inclusive:exclusive]
	// which outputs a slice --> references part of the existing array (a mutable view over the underlying array)
	// changing a value in (or appending to) `l` changes also `s`.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// from 0 (inclusive) to 4
	l = s[:5]
	fmt.Println("sl2:", l)

	// from 2 (inclusive) to the end
	l = s[2:]
	fmt.Println("sl3:", l)

	// Decl + init in 1 line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	// The `slices` package includes lots of utility functions
	// Here, `Equal` does shallow copy! Even though slices t and t2 do not reference the same underlying data array,
	// I think the individual string references reference the SAME string value (ie, "g", "h", etc.)
	// because of Go's string interning (immutable string values being shared).
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// In multi-dimensional slices, inner slices' length can vary!
	// It's like a vector of vectors
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	// [[0] [1 2] [2 3 4]]
	fmt.Println("2d: ", twoD)
}
