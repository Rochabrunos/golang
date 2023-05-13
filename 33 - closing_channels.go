package main

import "fmt"

/*
Closing a channel indicates that no more value will be sent on it
This can be useful to communicate completition to the channel's service
*/
func main() {
	// communicate the work to be done
	jobs := make(chan int, 5)
	// comminicate when the work is all done
	done := make(chan bool)

	go func() {
		for {
			j, more := <- jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	// close the jobs channel
	close(jobs)
	fmt.Println("sent all jobs")

	//blocks the excution until we receive a notification from the worker on channel
	<- done
}