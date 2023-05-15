package main

import (
	"fmt"
	"os"
	"bufio"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	str := "First things first!"
	d1 := []byte(str)
	err := os.WriteFile("../tmp/dat1.txt", d1, 0644)
	check(err)

	f, err := os.Create("../tmp/dat2.txt")
	check(err)

	defer f.Close()

	d2 := []byte(str)
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	//we had modified the file, WriteString just append
	n3, err := f.WriteString("\n1 - First things first!\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	
	//flushes writes to stable storage
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	//ensure that all buffered operations have been applied	
	w.Flush()


}