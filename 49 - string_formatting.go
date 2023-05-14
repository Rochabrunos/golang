package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1,2}
	//verbs: designed to formar general values
	fmt.Printf("struct1: %v\n", p)
	//+verb will include the struct's field names
	fmt.Printf("struct2: %+v\n", p)
	//# print a go syntax representation of the value
	fmt.Printf("struct3: %#v\n", p)

	//type: print the type of value
	fmt.Printf("type: %T\n", p)
	//formatting a boolean is pretty straightforward
	fmt.Printf("bool: %t\n", true)
	//base-10 formmating
	fmt.Printf("int: %d\n", 123)
	//binary representation
	fmt.Printf("bin: %b\n", 14)
	//print the character corresponding to the given integer
	fmt.Printf("char: %c\n", 64)
	//hex encoding
	fmt.Printf("hex: %x\n", 456)
	//decimal formatting
	fmt.Printf("float1: %f\n", 78.9)
	//scientific notation
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float2: %E\n", 123400000.0)
	//string basic print
	fmt.Printf("str1: %s\n", "\"string\"")
	//double-quote strings
	fmt.Printf("str2: %q\n", "\"string\"")
	//print pointer representation
	fmt.Printf("pointer: %p\n", &p)
	/* 
	When formatiing numbers we will often want to control the width and precision
	of the resulting figure.
	*/
	//to specify a width of an integer, use a number after the % in the verb
	fmt.Printf("width1:|%6d|%6d|\n7", 12, 435)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	//to left-justify
	fmt.Printf("width3: |%-6s|%-6s|\n", "foo", "bar")
	
	//sprintf format and returns a string without printiting it anywhere
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io:an %s\n", "error")
}