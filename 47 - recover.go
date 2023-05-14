package main

import "fmt"
/*
THIS CODE WILL NOT WORK PROPERLY: NEVER WILL PRINT "After mayPanic()""
*/
func mayPanic() {
	panic("a problem")
}

func main() {
	//recover has to be called with a deferred function
	//when the function panic defer will activate
	defer func() {
		//a recover call will catch the panic
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	
	//Will always panic and stops the execution of the main goroutine here
	mayPanic()

	fmt.Println("After mayPanic()")
}