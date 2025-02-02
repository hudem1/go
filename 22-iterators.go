// All types and functions within package `main` are visible in all files within the same package
// Therefore, we can reuse previous types/methods/functions/etc here.
//
// Iterators are lazy, meaning the values being iterated over are lazily evaluated:
// --> they are computed one by one at a time when needed instead of all upfront (eager evaluation)
//
// Lazy evaluation advantages:
// - memory-efficiency: avoid storing large data structure
// - infinite sequence: lazy eval is essential since we cannot precompute an infinity of values
// - on-demand exec: computes the next value only when requested (request by the `yield` fct), avoiding unnecessary computations
//
// Eager evaluation advantages:
// - random-access: provides O(1) access and does not require computing previous values like lazy eval
// - multiple iterations: precomputed array avoids redundant recomputation
// - simpler implementation: if computation is cheap & dataset is small, an eager approach might be simpler and faster
package main

import (
	"fmt"
	"iter"
	"slices"
)

// `All` returns an iterator, which is function with a special signature: `func(yield func(T) bool)`
// The iterator function takes another function as a parameter, called yield by convention (but the name can be arbitrary).
// It will call yield for every element we want to iterate over, and note yield’s return value for a potential early termination (if `false`).
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {

		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Iteration doesn’t require an underlying data structure, and doesn’t even have to be finite!
// Here’s a function returning an iterator over Fibonacci numbers: it keeps running as long as `yield` keeps returning true.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func iterators_main() {
	// Iterator usage outside of `range` loop with explicit `yield` fct
	fib := genFib()
	// Call the iterator function with a custom `yield` fct
	fib(func(n int) bool {
		fmt.Println(n)
		return n < 100 // Keep yielding while n < 100
	})

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// Since List.All returns an iterator, we can use it in a regular range loop.
	//
	// Note when using range, the `yield` callback is not explicitly provided, rather
	// Go internally provides it to the iterator returned by `lst.All`.
	//
	// However that being said, the body of the `range` block controls the returned boolean of `yield`:
	// - if the body hits `break`, the `yield` will return `false` ending the iteration.
	// - otherwise, the `yield` always returns true, and the iteration needs to be stopped within
	//   the iterator itself -> which is the case here below, where the iteration stops when the
	//   data structure `lst` does not have anymore elements to iterate over.
	for e := range lst.All() {
		fmt.Println(e)
	}

	// Packages like `slices` have a number of useful functions to work with iterators.
	// For example, `Collect` takes any iterator and collects all its values into a slice.
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	// Once the loop hits `break`, the `yield` fct passed to the iterator will return `false` ending the iteration.
	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
