package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	// "time"
)

/*
The channel-based approach, for synchronize access to shared 
state across multiple goroutines, aligns with Go's ideas of sharing
memory by communicating and having each piece of data owned by exacly
one goroutine

each go routine will send messages to the owning gorouting 
and recieve correspondent replies
*/

type readOp struct {
	key int
	resp chan int
}
type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {
	// these var accounts to the total operations we perform
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)
	// Here we have the state of our application
	go func() {
		var state = make(map[int]int)
		//keep waiting for the request
		for {
			select {
			case read := <- reads:
				read.resp <- state[read.key]
			case write := <- writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				//created the struture for the request
				read := readOp {
					key: rand.Intn(5),
					resp: make(chan int),
				}
				//send the request
				reads <- read
				//wait until the request is completed
				<- read.resp
				atomic.AddUint64(&readOps, 1)
				// time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				//create the struture for the write request 
				write := writeOp {
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool),
				}
				//send the request
				writes <- write
				//wait for answer
				<- write.resp
				atomic.AddUint64(&writeOps, 1)
				// time.Sleep(time.Millisecond)
			}
		}()
	}

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}