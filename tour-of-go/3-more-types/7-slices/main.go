package main

import "fmt"

// []T is a slice of items of type T
// A slice is a flexible, dynamic view into elements of an array
// create slice with elements 1-3 of "a": a[1:4]


func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Create slice from primes array
	var s []int = primes[1:4]

	fmt.Println(s)
}
