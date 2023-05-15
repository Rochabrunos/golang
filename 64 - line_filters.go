package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//wraps unbuffered os.Stdin with a buffered scanner
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(); scanner.Text() != "a"; scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	/*
	It is part of the example application that's why I didn't erase
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	*/
}