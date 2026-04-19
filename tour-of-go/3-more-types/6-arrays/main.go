package main

import "fmt"

// [n]T is an array of n items of type T
// An array's length is part of its type


func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)



	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
