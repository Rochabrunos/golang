package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	- Timers represent a single event in the future
	- We tell the time how long we want to wait
	- It provides a channel that will be notified at that time
	*/

	timer1 := time.NewTimer(2*time.Second)

	<- timer1.C
	fmt.Println("Timer 1 fired")

	// A useful feature of Timer is you can cancel the time before it fires
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired ")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2*time.Second)

}