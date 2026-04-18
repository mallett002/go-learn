package main

import "fmt"

// pointer type "*T"
// get memory address of var: "&"
// dereference pointer: "*" (go to value at memory address)


func main() {
	i, j := 42, 2701

	// set pointer to i
	p := &i

	// read i through pointer
	fmt.Println(*p)

	// set i through pointer
	*p = 21
	fmt.Println(i)

	// point to j
	p = &j

	fmt.Printf("memory address of j: %v\n", p)

	// divide j through pointer
	*p = *p / 37 // "val @ address = val @ address / 37"
	fmt.Printf("new val of j: %v\n", j)
}
