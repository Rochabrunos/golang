package main

import "fmt"

/*
	Pointers allows us to pass references to values and record within your program
*/

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	//deference
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial val:", i)
	fmt.Println("pointer addr:", &i)

	zeroval(i)
	fmt.Println("zeroval:", i)
	
	zeroptr(&i)
	fmt.Println("zeroptr", i)

}