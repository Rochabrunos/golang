package main

import (
	"flag"
	"fmt"
)

func main() {
	//(name, defaultValue, helper) (*string)
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	//another way to declare flag
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	//execute command line parsing
	flag.Parse()
	fmt.Println("word:",*wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}