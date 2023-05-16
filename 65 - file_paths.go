package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

/*
Filepath package providess functions to parse and construct
filepaths in a way that is portable between operating systems
*/
func main() {
	//Join should be used to construct path in a portable way
	p := filepath.Join("dir1", "dir2", "filepath")
	fmt.Println("p:", p)

	//removes superfuluous separators and directory changes
	fmt.Println("p2:",filepath.Join("dir1//", "filename"))
	fmt.Println("p3:", filepath.Join("dir1/../dir1", "filename"))

	//Dir and Base can be used to split a path to the directory and the file
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	//check if a path is absolute
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("C:/dir/file"))

	filename := "config.json"

	//extract the extension of the file
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	//find the file's name with extension removed
	fmt.Println(strings.TrimSuffix(filename, ext))

	//find a relative path between a "base" and a "target"
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}