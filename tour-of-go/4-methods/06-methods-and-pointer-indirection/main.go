package main

import "fmt"


type Vertex struct {
	X, Y float64
}

// Pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Pointer arg
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
    // vertex as a non-pointer value
    v := Vertex{3, 4}
    // with pointer receivers
    // doesn't need to be &v.Scale(2)
    // Go handles for you as convenience
    v.Scale(2)
    // does need to be &v here (pointer arg)
    ScaleFunc(&v, 10)

    // vertex as a pointer value
    p := &Vertex{4, 3}
    // works with pointers also
    p.Scale(3)
    ScaleFunc(p, 8)

    // print field names as well: %+v
    fmt.Printf("v: %+v; p: %+v \n", v, p)
}
