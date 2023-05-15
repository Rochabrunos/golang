package main

import (
	"fmt"
	"math/rand"
	"time"
)
/*
Note that this is not safe to use for random numbers you intend to be secret;
use crypto/rand for those.
*/
func main() {
	//returns a random int n, 0 <= n <= 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	//retruns a float64 f, 0.0 <= < 1.0
	fmt.Println(rand.Float64())
	// 5.0 <= f < 10.0
	fmt.Print((rand.Float64()*5) + 5, ",")
	fmt.Print((rand.Float64() * 5) + 5 )
	fmt.Println()

	/*
	default number generator is deterministic
	to produce varying sequences, give it a seed that changes
	*/
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	//using the same seed we going to generate a same sequence of pseudo-random numbers
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
	fmt.Println()
}