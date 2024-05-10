package main

import (
	"fmt"
)

func main() {
	runMultipleResults()
	runVarsWithInit()
}

// ------multipleResults---------------
func runMultipleResults() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func swap(x, y string) (string, string) {
	return y, x
}


// ------Run vars with init---------------
var i, j int = 1, 2

func runVarsWithInit() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}