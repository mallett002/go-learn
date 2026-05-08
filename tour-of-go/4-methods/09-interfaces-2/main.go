package main

import (
	"fmt"
)


// Any type that has a  milesLeft satisfies this interface
type Engine interface {
	milesLeft() uint8
}

type GasEngine struct {
	mpg     uint8
	gallons uint8
}

type ElectricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// Give GasEngine the milesLeft method (Implement the interface)
func (e GasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

// Give ElectricEngine the milesLeft method (Implement the interface)
func (engine ElectricEngine) milesLeft() uint8 {
	return engine.kwh * engine.mpkwh
}

// Accepts engine interface to handle both GasEngine and ElectricEngine
func canMakeIt(e Engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("Can make it")
	} else {
		fmt.Println("Need fuel")
	}
}

func main() {
	var myGasEngine GasEngine = GasEngine{mpg: 24, gallons: 15}
	fmt.Printf("Miles left: %v", myGasEngine.milesLeft())

	var myElectricEngine ElectricEngine = ElectricEngine{25, 15}
	fmt.Printf("Miles left: %v", myElectricEngine.milesLeft())

	canMakeIt(myElectricEngine, 50)
	canMakeIt(myGasEngine, 50)
}

