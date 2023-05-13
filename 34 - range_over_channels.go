package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	/* 
	- iterate over the values recieved from a channel
	- because we closed the channel above, the iteration terminates after receiving the 2 elements
	- for the case on we don't close the channel will raise an deadlock error
	*/
	for elem := range queue {
		fmt.Println(elem)
	}
}