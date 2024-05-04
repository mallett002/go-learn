package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type contactInfo struct {
	Name string
	Email string
}


type purchaseInfo struct {
	Name string
	Price float32
	Amount int
}



func main() {
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n %+v", contacts) // "%+v" adds the key to the object being printed. if just %v, it's just the values

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n %+v", purchases)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	var data, _ = os.ReadFile(filePath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}