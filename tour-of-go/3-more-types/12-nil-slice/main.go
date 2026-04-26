package main

import "fmt"

// zero value of slice is nil
// nil slice has len and cap of 0 and no underlying array

func main() {
	var s []int

	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("s is nil!")
	}
}
