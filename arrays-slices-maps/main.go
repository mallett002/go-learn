package main

import (
	"fmt"
	"time"
)

func main() {
	// arrayBasics()
	// sliceBasics()
	// mapBasics()
	speedTest()
}

func arrayBasics() {
	fmt.Println("arraay basics: ")

	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(intArr)

	// stored in contiguous memory locations
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
	fmt.Println()

	// immediately initialize arr:
	var intArr2 [3]int32 = [3]int32{1, 2, 3}

	intArrInferred := [3]int32{1, 2, 3}

	inferred := [...]int32{1, 2, 3}

	fmt.Println(intArr2)
	fmt.Println(intArrInferred)
	fmt.Println(inferred)
}

func sliceBasics() {
	fmt.Println("slice basics: ")
	// omit length val, we now have a slice:
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)

	// add values to slice:
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	// add multiple values
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// create slice using make func:
	var intSlice3 []int32 = make([]int32, 3, 8) // type, length, capacity
	fmt.Println(intSlice3)

	// looping over slice
	for i, v := range intSlice2 {
		fmt.Printf("Index: %v, Value: %v", i, v)
	}

	// while loop in go
	i := 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	// or
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// basic math:
	i--
	i++
	i += 10
	i -= 10
	i *= 10
	i /= 10

}

func mapBasics() {
	fmt.Println("map basics:")

	// make map:
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	// or with initial values
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}

	// get age of adam:
	fmt.Println(myMap2["Adam"])

	// returns boolean if exists
	var age, ok = myMap2["Adam"]
	fmt.Printf("age: %v, exists: %v\n", age, ok)

	age, ok = myMap2["adam"]
	fmt.Printf("age: %v, exists: %v\n", age, ok)

	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	// delete value from map
	delete(myMap2, "Adam") // no return value

	// looping over maps:
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	// get names and values
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}
}

func speedTest() {
	var capacity int = 1_000_000
	var testSlice = []int{}
	var testSliceWithCapacity = make([]int, 0, capacity)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, capacity))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSliceWithCapacity, capacity))
}

func timeLoop(slice []int, n int) time.Duration {
	var start = time.Now()

	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(start)
}
