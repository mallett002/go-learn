package main

import (
	"fmt"
)

func channel(outChan chan<- string, msg string) {
	outChan <- msg
}

func transferChan(outChan <-chan string, inChan chan<- string) {
	msg := <-outChan
	inChan <- msg
}

func main() {
	outChan := make(chan string, 1)
	inChan := make(chan string, 1)

	channel(outChan, "Howdy doodie")
	transferChan(outChan, inChan)

	fmt.Println(<-inChan)
}
