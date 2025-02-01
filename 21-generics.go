package main

import "fmt"

// Can use generics on functions
//
// An actual function in the `slices` package --> it takes a slice of any comparable type + an element of that type
// |--> the comparable allows to use `==` and `!=` on that type
//
// Special note: the `~` means the function allows custom slice types (like `type MyInts []int`),
// of course as long as the underlying type is comparable
// |--> in other words, it allows []E OR any named type whose underlying type is []E
// Does not work by default (without `~`) because a named type is a distinct type from the underlying []E
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// Can use generics on types
//
// Linked list
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Can use generics on methods
//
// Elements are pushed at the tail of the linked list
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// Returns all list elements as a slice
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		// Fyi, `append` takes `e.val` by value, so, it creates a copy of each val!
		// But it does a shallow copy (if T is of primitive type or struct without pointer, it does not matter, it copies the whole data)
		// If T is a reference type (*struct, slice, map, etc.) then only the pointer is copied, not the actual underlying data
		elems = append(elems, e.val)
	}
	return elems
}

func generics_main() {
	var s = []string{"foo", "bar", "zoo"}

	// Do not have to specify generic params, compiler infers them
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// Can still specify generic params explicitly if wished
	_ = SlicesIndex[[]string, string](s, "zoo")

	type MyInts []int
	var ints = MyInts{1, 2, 3}
	// Below line would not compile if `SlicesIndex` did not have `~` in its generic param
	fmt.Println("index of 3:", SlicesIndex(ints, 3))

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
