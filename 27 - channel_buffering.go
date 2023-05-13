package main

import "fmt"

/*
Buffered channels accepts a limited number of values without a corresponding receiever for those values
*/
func main() {
	messages := make(chan string, 3)
	
	messages <- "buffered"
	messages <- "channel"
	
	
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}