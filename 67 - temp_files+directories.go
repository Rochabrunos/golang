package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	
	//creates a file and open it for reading and writing
	f, err := os.CreateTemp("", "sample") //the first argument is the path
	check(err)

	fmt.Println("Temp file name:", f.Name())

	//ensure that the file will be deleted
	defer func() {
		//THE  FILE MUST BE CLOSED BEFORE DELETE IT
		f.Close()
		os.Remove(f.Name()) } ()

	_, err = f.Write([]byte{1,2,3,4})
	check(err)

	//creates a temporary directory
	dname, err := os.MkdirTemp("", "sampleDir")
	check(err)
	fmt.Println("Temp name dir:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	fmt.Println("fname:",fname) 
	err = os.WriteFile(fname, []byte{1,2}, 0666)
	check(err)


}	