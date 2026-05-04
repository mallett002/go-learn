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

		return last
	}
}

// actual solution
func fibonacci2() func() int {
	var curr, next = 0, 1

	return func() (fib int) {
		fib, curr, next = curr, next, curr+next

		// Scenario to help understand
		// curr: 1
		// next : 2

		// fib = curr (1)
		// curr = next (2)
		// next = curr + next (3)
		// return fib

		// visualizing this:
		// z: 0, x: 0, y: 1    (curr)
		// z: 0, x: 1, y: 1    (assigned)

		// z: 0, x: 1, y: 1    (curr)
		// z: 1, x: 1, y: 2    (assigned)

		// z: 1, x: 1, y: 2    (curr)
		// z: 1, x: 2, y: 3    (assigned)

		// z: 1, x: 2, y: 3    (curr)
		// z: 2, x: 3, y: 5    (assigned)

		// ...

		return
	}
}

func main() {
	// ex: 0,     1,     ?
	//    prev  curr    next

	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
