package main

import "fmt"

//pings only accepsts a channel for sending values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//ping only recieve values
//pongs only accepts a channel to send values
func pong(pings <-chan string, pongs chan<- string) {
	msg := <- pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passes message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}