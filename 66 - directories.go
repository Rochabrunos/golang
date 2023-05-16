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

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
func main() {
	err := os.Mkdir("subdir", 0755)
	check(err)

	defer func() { os.RemoveAll("subdir") } ()

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")
	//create a hieracy of directories
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	//get the directory's content, return an slice of os.DirEntry objects
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// err = os.Chdir("../../")
	// check(err)

	//visit recursively a directory
	fmt.Println("Visiting a directory")
	//Walk accepts a callback function to handle every file or directory visited
	err = filepath.Walk("subdir", visit)
	check(err)
}