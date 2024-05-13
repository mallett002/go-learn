package main

import (
	"fmt"
)

func main() {
	pointers()
}

func pointers() {
	pointerBasics()
}

func pointerBasics() {
	i, j := 42, 2701

	p := &i 							// point to i
	fmt.Printf("\n*p is: %v", *p)		// read i through the pointer
	fmt.Printf("\np is %v: ", p)		// print memory address of i
	*p = 21								// set i through the pointer
	fmt.Printf("\ni is: %v\n", i)

	fmt.Println("")

	fmt.Printf("j is: %v\n", j)
	p = &j								// point to j
	*p = *p / 37						// divide j through the pointer
	fmt.Printf("j is now: %v", j)
}