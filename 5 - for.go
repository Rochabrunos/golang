package main

import "fmt"

func main() {
	// Remember that's same thing as to explictly declare " var i integer = 1 "
	i:= 1

	// Very basic type, with a single condition
	for i < 3 {
		fmt.Println("Now i value is", i)
		i = i + 1
	}

	// A classic initial/condition/after for loop
	for j := 1 ; j <= 3 ; j++ {
		fmt.Println("Now j value is:", j);
	}

	// Loop without condition is allowed
	for {
		fmt.Println("loop")
		break //Or return must exist to avoid infinity loop
	}

	// 'continue' feature to jump to next iteration
	for n := 0; n <= 5 ; n++ {
		if n%2 == 0 {
			continue
		}
		// Just show the odd numbers
		fmt.Println(n)
	}


}
