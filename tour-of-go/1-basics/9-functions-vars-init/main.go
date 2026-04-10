package main

import "fmt"

// globel level vars with initializer
var i,j int = 1, 2

// function level vars
func main() {
	// if initializer present, can omit type (takes on type of initializer)
	var c, python, java = true, false, "No!"

	fmt.Println(i, j, c, python, java) // 0 false false false
}


