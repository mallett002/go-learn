package main

import "fmt"

// stringer interface (defined in fmt package)
// type Stringer interface {
//     String() string
// }
// fmt package, and many others look for this String() method to print values

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
    a := Person{"Arthur Dent", 42}
    z := Person{"Zaphod Beeblebrox", 9001}

    fmt.Println(a, z)
}
