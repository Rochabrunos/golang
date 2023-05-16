package main

import (
	"fmt"
	"os"
	"strings"
)
/*
Environment variables are a universal mechanism for conveying
configuration information to Unix programs
*/

func main() {
	//set key/value pair
	os.Setenv("FOO", "1")
	//get a valur for a key
	fmt.Println("FOO:", os.Getenv("FOO"))
	//return an empty string if key isn't present
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	//Environ: list all key/pair values KEY=value
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], "=", pair[1])
	}
}