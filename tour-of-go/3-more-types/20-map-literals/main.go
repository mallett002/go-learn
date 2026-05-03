package main

import "fmt"

// map literal keys are required

type Vertex struct {
	Lat, Lng float64
}

var m = map[string]Vertex {
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex {
		37.42202, -122.08408,
	},
}

// can omit type name inside literal
var other = map[string]Vertex {
	"Bell Labs": { // here
		40.68433, -74.39967,
	},
	"Google": { // and here
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
	fmt.Println(other)
}
