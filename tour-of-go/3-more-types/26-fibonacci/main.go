package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last := 0
	next := 1

	return func() int {
		nextNext := last + next	  	

		last = next
		next = nextNext

		return nextNext
	}
}

func fibonacci2() func() int {
	var last, next = 0, 1

	return func() (nextNext int) {
		nextNext, last, next = last, next, last+next

		return
	}
}

func main() {
	// ex: 0, 1, 1, 2, 3, 5, 8

	f := fibonacci2()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
