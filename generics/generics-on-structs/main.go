package main

import (
	"fmt"
)

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

// generics on this struct:
type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	// use the generic type to create a gasEngine object:
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg:     40,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh: 57.5,
			mpkwh: 4.17,
		},
	}

	fmt.Printf("\n gas car: %+v", gasCar)
	fmt.Printf("\n electric car: %+v", electricCar)
}
