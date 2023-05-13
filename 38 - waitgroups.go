package main

import (
	"fmt"
	"sync"
	"time"
)
/*
The WaitGroup is used to wait for all the goroutines lauched here to finish
*/

func worker(id int) {
	fmt.Printf("Worker %d\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func main() {

	var wg sync.WaitGroup
	//launch several goroutines and increment the WaitGroup counter for each
	for i := 1; i <= 5; i++ {
		//Add 1 to the counter of open jobs
		wg.Add(1)
		//avoid to reuse the same i value in each Goroutine closure https://go.dev/doc/faq#closures_and_goroutines
		i := i
		go func() {
			//defer the execution until the function doesn't return something 
			//Done() decrement the counter of open jobs
			defer wg.Done()
			worker(i)
		}()
	}
	//Wait until the counter of open jobs goes back to 0
	wg.Wait()
}