package main

import (
	"fmt"
	"time"
)

/*
Rate limiting
- controls resource utilization
- maintaining quality of service
- Go supports rate limiting with goroutines, channels and tickers
*/
func main() {
	limitTicker := 200 * time.Millisecond
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//this limiter channel will receive a value every 200 ms
	limiter := time.Tick(limitTicker)

	// limits to 1 request every limitTicker  
	for req := range requests {
		<- limiter // wait until receive a value
		fmt.Println("request", req, time.Now())
	}

	// this will allow burst of up to 3 events
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3 ; i++ {
		burstyLimiter <- time.Now()
	}

	// every 200 ms try to add an work
	go func(){
		for t := range time.Tick(limitTicker) {
			burstyLimiter <- t
		}
	}()

	//try to launch 10 more jobs, only the three first will benifit from burst capability
	burstyRequests := make(chan int, 10)
	for i := 3; i <=12; i++ {
		burstyRequests <- i
	}
	//close the channel
	close(burstyRequests)
	
	//get the results
	for req := range burstyRequests {
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}