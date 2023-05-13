package main

import (
	"fmt"
	"unicode/utf8"
)

/*
 -> A Go string is a read-olny slice of bytes
 -> As standards the text is encoded as UTF-8
 -> In other languages, strings are made of "characters"
 -> The concept of character is called a rune
	-> It's an integer that represent a Unicode code point

More information about it 
* https://go.dev/blog/strings
* https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/
*/

func main() {
	// S is a string represent "hello" in the Thai language
	const s = "สวัสดี"

	// The result here is 18, since strings are equivalent to []bytes
	fmt.Println("Len:", len(s))
	
	// This loop generates the hex values of all the bytes that constitute the code points in s
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	// Result of the for loop above : e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5 
	fmt.Println()
	// Some Thai characters are represented by multiple UTF-8 code points
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	// The result here is 6

	

}