package main

import "fmt"

// Range iterates over elements in a variety of built-in data structures
func ranges_main() {
	nums := []int{2, 3, 4}
	sum := 0
	// Range over slices
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	array := [2]int{1, 2}
	// Range over arrays
	for i, v := range array {
		fmt.Printf("key: %d -> value: %d\n", i, v)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	// Range over maps
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Range over strings
	// i: index in string
	// c: character unicode, ie a -> 97, b -> 98, etc.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
