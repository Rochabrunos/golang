package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


func main() {
	//we use an unsigned integer to represent our counter
	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops:", ops)
}