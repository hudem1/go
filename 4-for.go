package main

import "fmt"

func for_main() {
	// like a while
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic for loop
	// have to use ":=" shorthand, cannot use "var j = 0"
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// for loop over a range
	for i := range 3 {
		fmt.Println("range", i)
	}

	// "loop" in Cairo (infinite loop)
	// get out using break or return
	for {
		fmt.Println("loop")
		break
	}

	// can use "continue"
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
