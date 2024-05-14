package main

import (
	"fmt"
)

func main() {
	// pointers()
	// basicSlices()
	// maps()
	closures()
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

func basicSlices() {
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

	// When update slice, updates *names:
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	fmt.Println(cap(b))

	// Extending slice:----------------------------
	s := []int{2, 3, 5, 7, 11, 13}

	// slice it to give it 0 length
	s = s[:0]
	printSlice(s)

	// extend its length
	s = s[:4]
	printSlice(s)

	// drop its first 2 values
	s = s[2:]
	printSlice(s)


	// nil slice has no length nor cap and no underlying array
	var nilSlice []int
	printSlice(nilSlice)
	if nilSlice == nil {
		fmt.Println("nil!")
	}

	// Making slices with make func-------------------------
	// make(type, length, capacity)
	makeSliceA := make([]int, 5)
	printSlice(makeSliceA)
	
	makeSliceB := make([]int, 0, 5)
	printSlice(makeSliceB)

	makeSliceC := makeSliceB[:2]
	printSlice(makeSliceC)

	makeSliceD := makeSliceC[2:5]
	printSlice(makeSliceD)

	// appending to the slice
	makeSliceA = append(makeSliceA, 1)
	printSlice(makeSliceA)

	makeSliceB = append(makeSliceB, 1, 2, 3, 4, 5, 6)
	printSlice(makeSliceB)

	// looping over a slice with "range"
	for i, v := range makeSliceB {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func printSlice(slice []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
}

func maps() {
	var stuff map[string]int = make(map[string]int)

	stuff["A"] = 1
	fmt.Println(stuff["A"])

	m := make(map[string]int)
	m["answer"] = 42
	fmt.Println(m)

	delete(m, "answer")
	fmt.Println(m)

/*
	Insert or update an element in map m:
	m[key] = elem

	Retrieve an element:
	elem = m[key]

	Delete an element:
	delete(m, key)

	Test that a key is present with a two-value assignment:
	elem, ok = m[key]
*/
}

func closures() {
	// Go functions may be "closures"
	pos, neg := adder(), adder() // get the inner functions

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
		fmt.Println(neg(i * -2))
		fmt.Println("")
	}
}

// adder returns a func that takes in an int and returns an int:
func adder() func(int) int {
	sum := 0

	return func(x int) int {
		fmt.Println("sum is currently: ", sum)
		fmt.Println("x is currently:", x)

		sum += x

		return sum
	}
}