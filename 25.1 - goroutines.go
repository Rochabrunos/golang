package main

import (
	"fmt"
)

type Request struct {
	args	[]int
	f	func([]int) int
	resultChan	chan int
}

func sum(a []int) (s int) {
	for _, val := range a {
		s += val
	}
	return
}

func handle(queue chan *Request) {
	for req := range queue {
		req := req
		req.resultChan <- req.f(req.args)
	}
}

func client() {
	
}

func main() {
	//client req
	request := &Request{[]int{3,4,5}, sum, make(chan int, 1)}
	//send the request
	clientQueue := make(chan *Request, 1)
	clientQueue <- request
	close(clientQueue)
	//call hadle function on server
	handle(clientQueue)	
	//wait for answer
	fmt.Printf("answer: %d\n", <-request.resultChan)
}