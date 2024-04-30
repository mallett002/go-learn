package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// var mutex = sync.Mutex{} // mutual exclusion
var mutex = sync.RWMutex{} // Read/Write mutual exclusion
var wg = sync.WaitGroup{}
var dbData = []string{"1", "2", "3", "4", "5"}
var results = []string{}

func main() {
	t0 := time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1) // add to the counter
		go dbCall(i)
	}

	wg.Wait() // wait for the counter to go to zero (all go routines have finished)

	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	// fmt.Println("The result from db is: ", dbData[i])

	// mutex.Lock() // if already locked, waits until unlocked and then locks
	// results = append(results, dbData[i])
	// mutex.Unlock()
	save(dbData[i])
	log()

	wg.Done() // decrements the wg counter
}

func save(result string) {
	// mutex.Lock():
		// only proceeds if all, full and read, locks are cleared, else waits until cleared
		// vv Prevents writing data while others are reading/writing to the same data
	mutex.Lock()
	results = append(results, result)
	mutex.Unlock()
}

func log() {
	mutex.RLock() // read lock: Checks for full lock, "mutex.Lock()", b4 reading. Don't want to read while writing.
	fmt.Printf("\nThe current results are: %v", results)
	mutex.RUnlock()
}
