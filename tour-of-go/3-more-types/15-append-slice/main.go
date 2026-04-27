package main

 import (
	"fmt" 
)

// append adds items to a slice
// if underlying array too small to fit items, a new array is created to fit

func main() {
	var s []int
	printSlice("s", s) // []

	// append works on nil slices
	s = append(s, 0)
	printSlice("s", s) // [0]

	// the slice grows as needed
	s = append(s, 1)
	printSlice("s", s) // [0 1]

	// we can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice("s", s) // [0 1 2 3 4]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s, len=%d, cap=%d, %v\n",
		s, len(x), cap(x), x)
}

