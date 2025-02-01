package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// If a function has several consecutive parameters of the same type,
// we can omit the type up to the last param that declares the type in question
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Functions can return multiple values
// Often used for returning both result & error values
func vals() (int, int) {
	return 3, 7
}

// Variadic functions can be called with any number of trailing args
func sum(nums ...int) {
	// Print is actually a common variadic fct
	fmt.Println(nums, " ")

	total := 0
	// Within the fct, the type of `nums` is interpreted as a slice (here []int)
	// We can therefore do exactly the same as on slices (here, iterate over it with range)
	// Note: range returns index, value
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func functions_main() {
	// Basic functions
	res := plus(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 = ", res)

	// Multiple returns
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	// Variadic functions
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// Can "destructure" a slice into a list of elements (only in fct calls?)
	// But the following gives compilation error: test := nums...
	sum(nums...)
}
