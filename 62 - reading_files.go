package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//load the file's content into memory
	dat, err := os.ReadFile("../tmp/dat.txt")
	check(err)
	//print the content of bytes
	fmt.Println(dat)
	//coerce into string
	fmt.Println(string(dat))

	//os.Open() gives more control about what part of a file are read 
	f, err := os.Open("../tmp/dat.txt")
	check(err)

	defer func() {
		f.Close()
	}()

	//buffer
	b1 := make([]byte, 5)
	//number of bytes readed
	n1, err := f.Read(b1)
	check(err)
	fmt.Println("n1:", n1)
	fmt.Printf("%x bytes: %s\n", n1, string(b1[:n1]))

	// 0 means relative origion of the file; 1 to the current offset; and, 2 to the end
	o2, err := f.Seek(-6,2)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6,0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//rewind
	_, err = f.Seek(0,0)
	check(err)

	//bufio implements buffered reader
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
}

