package main
// Go will infer the type of initialized variables

import "fmt"

func main() {
	// "var" : Declares one or more variables
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// Short for '' var d bool = true '' 
	var d = true
	fmt.Println(d)

	// Variables without corresponding initialization are "zero-valued"
	var e int
	fmt.Println(e)

	// Short declaration for '' var f string = "apple" ''
	// This type of declaration is only available inside functions
	f:= "apple"
	fmt.Println(f)

}