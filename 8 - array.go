package main

import "fmt"

func main() {
	// Array with length of 5 integer elements
	var a [5]int
	// By default an arry is zero-valued like variables
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// buit-in function length len()
	fmt.Println("len:", len(a))

	// Declaration and inicialization
	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:", b)

	// Two dimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)
}