package main

import "fmt"


// types inferred:

func main() {
	i := 42 // int 	
	f := 3.1452 // float64
	g := 0.864 + 0.5i // complex128

	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}
