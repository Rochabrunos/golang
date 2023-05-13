package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	// sort is in-place, does not return a new one
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	//check if a slice are already in sorted order
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}