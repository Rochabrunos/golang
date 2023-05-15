package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!@$*&()'-=@~"

	//encoding requires a []byte
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	//decoding
	sDec, err := b64.StdEncoding.DecodeString(sEnc)
	if err != nil{
		panic(err)
	} 
	fmt.Println(string(sDec))

	//encode/decodes using URL-compatible base64 format
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}