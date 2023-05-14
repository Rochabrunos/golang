package main

import "os"

/*
Panic generally means something went unexpctedly wrong
*/
func main() {
	//we just use a panic to check for unexpected errors
	//exit with a non-zero
	panic("a problem")

	//comment the above line to run the code below
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}