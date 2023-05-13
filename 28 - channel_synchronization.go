package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

func main() {
	//the channel done will be used to notify that work is done
	done := make(chan bool, 1)
	
	fmt.Println("First move")
	go worker(done)
	fmt.Println("Last move")

	//blocks the excution until receive a notification from the worker on channel
	<- done
}