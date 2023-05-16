package main

import "embed"

/*
//go:embed is a compiler directive that allows programs
to include arbitrary files and folders in the Go binary
at build time
*/

//embeds the content of file in the string immediately following it
//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//we can also embed multiple files or even folders with wildcards
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS 

func main() {

	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/hash1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/hash2.hash")
	print(string(content2))
}