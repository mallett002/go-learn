package main

// go programs are comprised with packages
// this is the main package
// import packages
// imported package name is the last part of the path, ex: math/rand -> rand

import (
	"fmt"
	"math/rand"
)

func main () {
	fmt.Println("My favorite number is", rand.Intn(100))
}
