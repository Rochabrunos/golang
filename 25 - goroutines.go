package main

import (
	"fmt"
	"time"
)

/*
A goroutine is a lightweight thread execution
*/

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//running synchronously
	f("direct")

	//this will execute concurrently
	go f("goroutine")

	// time.Sleep(time.Second)

	//we can also start a goroutine for an anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}