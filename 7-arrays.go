package main

import "fmt"

func arrays_main() {

	// By default, an array is zero-valued.
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	// Decl + init in 1 line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Make compiler figure array size
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Specifying an index will fill previous indexes' value with 0
	b = [...]int{100, 3: 400, 500}
	// [100, 0, 0, 400, 500]
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
