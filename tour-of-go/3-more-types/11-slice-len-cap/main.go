package main

import "fmt"

// slices have length and capacity
// length is amt of items in the slice
// capacity is amt of items in underlying array starting with first item in slice

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// slice the slice to give it 0 length
	s = s[:0]
	printSlice(s)

	// extend its length - note the underlying array didn't change
	s = s[:4]
	printSlice(s)

	// drop the first 2 values
	s = s[2:]
	printSlice(s)

	fmt.Println("From array:")

	a := [6]int{2, 3, 5, 7, 11, 13}
	b := a[:]
	printSlice(b)
}

func printSlice(s []int) {
	fmt.Printf("len %d; cap %d slice: %v\n", len(s), cap(s), s)
}
