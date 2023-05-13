package main

// The same for import "fmt" \n import "math"
import (
	"fmt"
	"math"
)

// Declaring a contant; it can appear anywhere a var statement can
const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 50000

	// Constant expression perform arithmetic with arbitrary precision
	const d = 3e20 / n

	// A numeric constant has no type until it's explicit given one
	fmt.Println(d)
	fmt.Println(int64(d))

	// The type can be given by context, e.g., variable assignement or function call
	fmt.Println(math.Sin(n))
}