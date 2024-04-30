package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var dbData = []string{"1", "2", "3", "4", "5"}

func main() {
	t0 := time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1) // add to the counter
		go dbCall(i)
	}

	wg.Wait() // wait for the counter to go to zero (all go routines have finished)

	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from db is: ", dbData[i])
	wg.Done() // decrements the wg counter
}
