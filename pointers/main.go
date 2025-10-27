package main

import "fmt"

func main() {
	// basics()
	// slices()
	functions()
}

func basics() {
	// pointers hold memory address
	var p *int32   // holds memory address of int32 (starts as nil here)
	p = new(int32) // now has memory location, with default value (0, "", false etc)

	// get value of a memory address (pointer)
	var val int32 = *p // de-referencing the pointer

	var i int32

	fmt.Printf("The value p points to is: %v\n", val)
	fmt.Printf("The value of i is: %v\n", i)

	// change value of a pointer (reference value of a pointer)
	*p = 10

	// create a pointer from address of another variable using "&"
	var newPointer *int32 = &i // get memory address, not its value

	// change newPointer, changes i as well (same place in memory)
	*newPointer = 1
	fmt.Printf("The value of i is: %v\n", i)

	// no pointers
	var k int32 = 2
	i = k // copy value of k into i's memory location
	fmt.Printf("The value of k is: %v\n", k)
	fmt.Printf("The value of i is: %v\n", i)
}

func slices() {
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice

	// changes both: (slices contain pointers to an underlying array under the hood)
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
}

func functions() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of thing1 is: %p", &thing1)

	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)

	// use pointer instead (uses same memory location)
	// uses the memory location (4 - 8 bits instead of complete copy of all the data)
	var resultWithPointer [5]float64 = squareWithPointer(&thing1)
	fmt.Printf("\nThe memory location of thing1 is: %p", &thing1)
	fmt.Printf("\nThe resultWithPointer is: %v", resultWithPointer)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of thing2 is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return thing2
}

func squareWithPointer(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of thing2 is: %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return *thing2
}
