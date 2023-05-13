package main

import "fmt"

/*
	Documentation: https://go.dev/blog/slices-intro
	What are the Slices?
	-> Slices are an importante data type in Go
	-> Give a more powerful interface to sequences than arrays
	-> Unlike arrays, slices are typed only by the elements the contains
*/
func main() {

	// make function: creates non-zero length slice (initially zero-valued)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// slice has the builtin function len() like arrays
	fmt.Println("len:", len(s))

	// append function: return a slice containing the new element
	s = append(s, "d")
	fmt.Println("apd:", s)
	// append can be create a new slice with multiples new elemtents
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	// copy function: make a copy of slice S into slice C
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice function: get a slice from the slice S from index 2 to 4
	l := s[2:5]
	fmt.Println("sl1:", l)

	// get a slice from the slice S from index 0 to 4
	l = s[:5]
	fmt.Println("sl2:", l)

	// get a slice from the slice S from index 2 to the last element
	l = s[2:]
	fmt.Println("sl3:", l)

	// Declaration with incialization
	t := []string{"g", "h", "i"}
	fmt.Println("dlc:", t)

	// It is also possible to make a multidimensional slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}