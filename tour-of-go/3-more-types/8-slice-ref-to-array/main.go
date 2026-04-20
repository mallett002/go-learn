package main

import "fmt"

// slice doesn't store any data
// it just describes a section of an underlying array
// changing items in slice changes underlying array
// other slices sharing same array will see those changes


func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// change slice (changes underlying array)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
