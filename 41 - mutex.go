package main

import (
	"fmt"
	"sync"
)

/*
Mutexes help us handle simultaneous data access in concurrent enviroment
*/

type Container struct {
	mu sync.Mutex
	counters map[string]int
}

//mutexes must not be copied, use always pointers
func (c *Container) inc(name string){
	c.mu.Lock() //lock mutex
	defer c.mu.Unlock() //defer until the end of function
	c.counters[name]++
}

func main() {
	c := Container {
		counters: map[string]int{"a":0, "b":0},
	}
	
	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n ; i++{
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)

}

