package main

import (
	"fmt"
	"math/rand"
	"sync" // Used to wait for threads to complete before continuing!
	"time"
)

var m = sync.Mutex{}      // Create mutex
var wg = sync.WaitGroup{} // Create wait group
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)    // Tell the wait group to add one to the counter
		go dbCall(i) // write go to use concurrency
	}
	wg.Wait() // Tell the program to wait untill the counter reaches 0
	fmt.Printf("Total executiontime: %v\n", time.Since(t0))

	// Do the same, but this time, append the result to a global slice
	t0 = time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbWrite(i)
	}
	wg.Wait()
	fmt.Printf("The writing took %v seconds\n", time.Since(t0))
	fmt.Printf("The results are %v\n", results)

}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	wg.Done() // Tell the wait group to subtract one from the counter
}

func dbWrite(i int) {
	time.Sleep(time.Duration(2000) * time.Millisecond)
	m.Lock() // Lock mutex to ensure safe writing!
	results = append(results, dbData[i])
	m.Unlock() // Unlock to allow others to write.
	wg.Done()
}
