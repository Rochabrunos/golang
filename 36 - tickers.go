package main

import (
	"fmt"
	"time"
)

/*
	- Tickers are for when you want to do something repeatedly at regular intervals
*/
func main() {
	//A channel is created for sent values
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <- done:
				return
			case t := <- ticker.C:
				fmt.Println("Ticker at", t)
			}
		}
	}()

	time.Sleep(2000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}