package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	//define an flag (required)
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "bool enable")
	fooName := fooCmd.String("name", "", "string name")

	barCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "int level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("	enable:", *fooEnable)
		fmt.Println("	name:", *fooName)
		fmt.Println("	tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("	level:", *barLevel)
		fmt.Println("	tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}