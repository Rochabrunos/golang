package main

import (
	"bytes"
	"fmt"
	"regexp"
)
/*
Stay hungry, see for more information: https://pkg.go.dev/regexp
*/
func main() {
	//test whether a pattern matches a string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println("regex:",r)

	//test whether a pattern matches a string
	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("punch punch"))
	
	//find the fist match and return the start and end indexes
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	//find the whole-pattern maches and the submatches 
	fmt.Println(r.FindStringSubmatch("peach punch"))
	//to return the indexes
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	//find all instances that matches
	fmt.Println(r.FindAllString("peach punch inch", -1))

	//"All" variation applied to other functions
	//providing non-negative number limits the number of mathces
	fmt.Println("all:",r.FindAllStringSubmatchIndex("peach punch inch", 2))

	//we can also provide other data types dropping "String" function name
	fmt.Println(r.FindSubmatchIndex([]byte("peach")))

	//using global value we can use MustCompile instead of Compile
	//MustCompile panics instead of return an error
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	//regexp can also be used to replace subsets of string
	fmt.Println(r.ReplaceAllString("a peach punch", "<fruit>"))

	//allows you to transform matched text with a given function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}