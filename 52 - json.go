package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page int
	Fruits []string
}
/*
-> only exported filds will be encoded/decoded in JSON
-> fields must always start with Capital letters to be exported
-> Here we go again, see more: https://go.dev/blog/json
*/
type response2 struct {
	Page int `json:"pageNumber"`
	Fruits []string `json:"fruitsSelling"`
}

func main() {

	//encoding basic data types to JSON strings
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	//working with slices and maps
	slcD := []string{"apple", "orange", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{
		"apple": 5,
		"lettuce": 7,
	}
	mapB, _ := json.Marshal(mapD)
	fmt.Printf("%q\n",string(mapB))

	/*
	The JSON package can automatically encode custom data types
	It will only include exported filds
	By default will use the field names as the JSON keys
	*/

	resp1D := &response1{
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	resp1B, _ := json.Marshal(resp1D)
	fmt.Println(string(resp1B))

	//It is possible to customize the fields as in respose2 struct
	resp2D := &response2 {
		Page: 1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	resp2B, _ := json.Marshal(resp2D)
	fmt.Println(string(resp2B))

	//feed some json data to decode
	byt := []byte(`{"pageNumber":1,"fruitsSelling":["apple","peach","pear"]}`)
	//To decode we need to provide a variabel where the JSON package can put the decoded data
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	} 
	fmt.Println(dat)
	num := dat["pageNumber"].(float64)
	fmt.Println(num)

	strs := dat["fruitsSelling"].([]interface{})
	fmt.Println(strs)
	//access nested data require a series of conversion
	strs1 := strs[0].(string)
	fmt.Println(strs1)

	//we can also decode JSON into custom data types
	res := response2{}
	json.Unmarshal(byt, &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// we can also stream JSON encoding
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(resp2D)
}