package main

import "fmt"

// make([]T, len, cap)
// creates a zeroed array and returns slice that refers to it

func main() {
	a := make([]int, 5)
	printSlice("a", a) // [0,0,0,0,0]

	// 0 len and 5 cap
	b := make([]int, 0, 5)
	printSlice("b", b) // []

	c := b[:2]
	printSlice("c", c) // [0,0]

	d := c[2:5] // indices 2, 3, 4
	printSlice("d", d) // [0,0,0]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s, len=%d, cap=%d, %v\n",
		s, len(x), cap(x), x)
}
