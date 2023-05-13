package main

import "fmt"

/*
	Range iterates over elements in a variety of data structure
	Range return each element with its index
	Range cannot guarantee the order which the elements is returned
*/
func main() {
	nums := []int{1,2,3,4}
	sum := 0

	for _,num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if(num == 3) {
			fmt.Println("index:", i)
		}
	}
	// *** sometimes the banana is printed prior than apple *** 
	kvs := map[string]string{"a": "apple", "b":"banana"}
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// For more information see on https://gobyexample.com/strings-and-runes
	for i,c := range "go" {
		fmt.Println(i, c)
	}
}