package main

import "fmt"

// slice literal is like array literal without length
// array literal: [6]int{2, 3, 5, 7, 11, 13}
// slice literal: []int{2, 3, 5, 7, 11, 13} creates same array as above, and a slice to reference it

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, false}
	fmt.Println(r)

	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
	}
	fmt.Println(s)
}
