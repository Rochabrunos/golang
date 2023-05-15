package main

import (
	"fmt"
	"strconv"
)
/*
The built-in 
*/

func main() {
	// this tells how many bits of precision to parse
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	//0 means the base of string
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	//ParseInt will recognize hex-formatted numbers
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	//ParseUint is also available
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	if _, e := strconv.Atoi("wat"); e != nil {
		panic(e)
	}
	
}