package main

import "fmt"

/*
	Go supports anonymous functions which can form closures
	Are useful when we need to define a function inline without having to name it
*/

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	
	newInts := intSeq()
	fmt.Println(newInts())
}