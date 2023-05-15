package main

import (
	"crypto/sha256"
	"fmt"
)

/*
SHA256 hashes are frequently used to compute short identities 
for binary or text blobs
https://en.wikipedia.org/wiki/Cryptographic_hash_function
*/

func main() {
	s := "sha256 this string"

	//start a new hash
	h := sha256.New()

	//Write method expects bytes... 
	//whether needed use []byte(string) to coerce it to bytes
	h.Write([]byte(s))

	//convert the slice of bytes into a human-reabable hex format
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Println(h)
	fmt.Printf("%x\n", bs)
}