package main

import "fmt"

/*
	Veriadic functions can be called with any number of trailing arguments
	fmt.Println() is a common variadic function 
*/

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("tol:", total)
}
func main() {
	sum()
	sum(1,2)
	sum(1,2,3)

	nums := []int{1,2,3,4}
	sum(nums...)
}