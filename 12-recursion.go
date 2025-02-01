package main

import "fmt"

// Basic recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func recursion_main() {
	fmt.Println(fact(7))

	// Declare anonymous function
	var fib func(n int) int

	// - Initialize anonymous function
	// - Anonymous functions can also do recursion but they need
	//   to be declared before to be able to call themselves recursively
	// - If decl + init in 1 line, cannot do recursion because their name/var had not been declared yet
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
