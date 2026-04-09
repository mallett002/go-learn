package main

import "fmt"

// Func can have named return values
// Treated as vars at beginning of function
// Can just do empty return and will return them

// x: ( 10 * 4 = 40 ) -> ( 40 / 9 ) = 4.444 -> 4
// y: ( 10 - 4 ) -> 6
// ratio: 4:5 (x gets 4 parts and y gets 5 parts)
func createFourToFiveRatio(num int) (fourParts, fiveParts int) {
	totalParts := 9 // 4 + 5

	fourParts = num * 4 / totalParts
	fiveParts = num - fourParts // get the rest

	return
}

func main() {
	a, b := createFourToFiveRatio(10) // -> 4, 6
	fmt.Println(a, b)
}


