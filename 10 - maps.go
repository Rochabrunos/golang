package main

import "fmt"

/*
	Maps are associative date type, e.g., hashes and maps
*/
func main() {

	// make create an empty map: make(map[key-type]value-type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v2 := m["k3"]
	fmt.Println("v2:", v2)

	fmt.Println("len:", len(m))

	// builtin function delete removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("del:", m)

	// the optional second element indicates if the key is present in the map
	_, prs := m["k2"]
	fmt.Println("prs_k2:", prs)

	value, prs := m["k1"]
	fmt.Println("prs_k1:", value, prs)

	// declaration and inicialization
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}