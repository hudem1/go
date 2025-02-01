// Struct embedding: embed a struct within another struct
package main

import "fmt"

type base struct {
	num int
}

// method on `base` struct
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// `container` struct embeds `base` struct
//   - The embedded struct needs to appear like below without a field name to be considered embedded,
//     just the type name by itself --> a field without a name
//
// If the embedded struct has methods, they also become methods of the embedding struct and can be called on it!
//
// Otherwise, if `container` had a field of type base `fieldName base`, the `base` struct would not be embedded
// and therefore methods from `base` are not inherited in `container`
type container struct {
	base
	str string
}

func struct_embedding_main() {
	// As we've seen in `structs` chapter, we cannot mix fields naming & values by themselves (literals) -> compilation error
	// and therefore, we can write either the notation either `co` (field naming) or `co2` (literals)
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	// or
	co2 := container{
		base{
			num: 2,
		},
		"some name",
	}
	fmt.Println(co2)

	// Can access embedded struct field directly on embedding struct `co`/`co2` as well as through embedded struct `base`
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)

	// Since `container` embeds `base`, the methods of `base` also become methods of a `container`.
	// Here we invoke a method that was embedded from `base` directly on `co`.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// `container` struct implements the `describer` interface because
	// `container` embeds `base` --> therefore `describe` method can be called on `container`
	var d describer = co
	fmt.Println("describer:", d.describe())
}
