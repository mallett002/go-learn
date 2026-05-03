package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// zero value of map is nil
// nil map has no keys
// nil map can't get keys added

// nil map:
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)

	m["will"] = Vertex{1.2, 2.2}	

	fmt.Println(m["will"])
}

