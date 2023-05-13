package main

import "fmt"

/*
Channels are the pipes that connect concurrent goroutines
You can share values between different channels
You can receive velues into another goroutine also

*/

func main() {
	//create a new channel with make(chan val-type)
	messages := make(chan string)

	//send value into a channel using channel <- syntax
	go func() {messages <- "ping"}()

	//recieving the value <- channel syntax
	msg := <- messages
	fmt.Println(msg)

	/*
	By default sends and recieves block until both are ready
	*/
}