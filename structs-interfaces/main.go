package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// ownerInfo owner
	// owner // owner: owner --> reference it with "myEngine.name"
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
	// var myEngine gasEngine
	var myEngine gasEngine = gasEngine{mpg: 24, gallons: 15}
	// var myEngine gasEngine = gasEngine{24, 15, owner{"Will"}}
	myEngine.mpg = 20

	// fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
	// Can just use ".name syntax" for owner now:
	// fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
	fmt.Println(myEngine.mpg, myEngine.gallons)

	// anonymous struct:
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	// calling milesLeft method on struct
	// fmt.Printf("Miles left: %v", myEngine.milesLeft())

	// using engine interface:
	var myElectricEngine electricEngine = electricEngine{25, 15}
	var myGasolineEngine gasEngine = gasEngine{29, 20}
	canMakeIt(myElectricEngine, 50)
	canMakeIt(myGasolineEngine, 50)
}

// methods on struct
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
