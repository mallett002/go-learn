package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// ownerInfo owner
	// owner // owner: owner --> reference it with "myGasEngine.name"
	// int  field type int with name type int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type engine interface {
	milesLeft() uint8
}

// type owner struct {
// 	name string
// }

func main() {
	// var myGasEngine gasEngine
	var myGasEngine gasEngine = gasEngine{mpg: 24, gallons: 15}
	// var myGasEngine gasEngine = gasEngine{24, 15, owner{"Will"}}
	myGasEngine.mpg = 20

	// fmt.Println(myGasEngine.mpg, myGasEngine.gallons, myGasEngine.ownerInfo.name)
	// Can just use ".name syntax" for owner now:
	// fmt.Println(myGasEngine.mpg, myGasEngine.gallons, myGasEngine.name)
	fmt.Println(myGasEngine.mpg, myGasEngine.gallons)

	// anonymous struct:
	var myGasEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Println(myGasEngine2.mpg, myGasEngine2.gallons)

	// calling milesLeft method on struct
	// fmt.Printf("Miles left: %v", myGasEngine.milesLeft())

	// using engine interface:
	var myElectricEngine electricEngine = electricEngine{25, 15}
	var myGasolineEngine gasEngine = gasEngine{29, 20}
	canMakeIt(myElectricEngine, 50)
	canMakeIt(myGasolineEngine, 50)
}

// Give gas & electric engines hte milesLeft method (Implement the interface)
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (engine electricEngine) milesLeft() uint8 {
	return engine.kwh * engine.mpkwh
}

// Accepts engine interface to handle both gasEngine and electricEngine
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("Can make it")
	} else {
		fmt.Println("Need fuel")
	}
}
