package main

import "fmt"

// Returns a function that itself returns an int
func intSeq() func() int {
	i := 0
	// Anonymous function
	return func() int {
		i++
		return i
	}
}

func closures_main() {
	// The function nextInt captures its own `i`, which will be updated
	// everytime we call `nextInt`
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	// The function nextInt captures its own `i` independent of nextInt's
	newInts := intSeq()
	fmt.Println(newInts()) // 1

	// Declare an anonymous function
	// See also next chapter `12-recursion`
	var myFunc = func() int {
		return 1
	}
	myFunc()
}
