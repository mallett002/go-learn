package main

import "fmt"

// when slicing you may omit lower and upper bounds to use the defaults

func main() {
	var nums [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// these are all the same:
	one := nums[0:10]
	two := nums[:10]
	three := nums[0:]
	four := nums[:]

	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(four)
}
