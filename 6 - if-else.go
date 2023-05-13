package main 

import "fmt"

func main() {
	// Classic example of if/else statement
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// else statement is optional
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// The context is inherent the 'if' statement
	if num:= 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// "num" variable doesn't exist here (out its of context)

	// The if no ternary if in Go
}