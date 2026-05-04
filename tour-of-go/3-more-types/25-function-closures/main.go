package main

import "fmt"

func adder() func(int) int {
	sum := 0

	// this closure bound to it's own sum function (doesn't get reset to 0 on each call)
	// if you keep using this function, it "remembers" what sum was on next call
	return func(x int) int {
		fmt.Printf("x: %d\nsum: %d\n", x, sum)

		sum += x

		return x
	}
}


func main() {
	posAdder := adder()
	negAdder := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			"res posAdder:", 
			posAdder(i),
		)

		fmt.Println(
			"res negAdder:", 
			negAdder(-2*i),
		)

		fmt.Println()
	}
}
