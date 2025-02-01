package main

import "fmt"

type person struct {
	name string
	age  int
}

// Go is a garbage collected language; you can safely return a pointer to a local variable.
// It will only be cleaned up by the garbage collector when there are no active references to it.
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

func structs_main() {
	// No need to name fields if declared in correct order
	fmt.Println(person{"Bob", 20})

	// Decl with fields naming
	fmt.Println(person{name: "Alice", age: 30})
	// Cannot mix fields naming & values by themselves -> compilation error
	// fmt.Println(person{name: "Alice", 30})

	// Can omit fields, which will be zero-valued
	fmt.Println(person{name: "Fred"})

	// Prints the address of the struct
	ann := person{name: "Ann", age: 40}
	fmt.Println(&ann)        // prints: &{Ann 40} --> Go dereferences struct pointers when printing them
	fmt.Printf("%p\n", &ann) // prints the actual address of ann

	// It's idiomatic to encapsulate new struct instance creation
	// in constructor functions
	fmt.Println(newPerson("John"))

	s := person{name: "Sean", age: 50}
	// Dot notation to access fields
	fmt.Println(s.name)

	sp := &s
	// Dot notation on a pointer automatically dereferences it (same as below)
	fmt.Println(sp.age)
	fmt.Println((*sp).age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// Anonymous struct type
	// If a struct type is only used for a single value we don't need to give it a name
	dog := struct {
		name   string
		isGood bool
	}{
		// Can also use fields naming (like for a normal struct)
		"Rex",
		true,
	}
	fmt.Println(dog)
}
