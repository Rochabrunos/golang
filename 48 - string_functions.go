package main

import (
	"fmt"
	str "strings"
)

var p = fmt.Println

func main() {
	p("Countains:	", str.Contains("test", "es"))
	p("Count:	", str.Count("test", "t"))
	p("HasPrefix:	", str.HasPrefix("teste", "te"))
	p("HasSuffix:	", str.HasSuffix("teste", "te"))
	p("Index:	", str.Index("teste", "t"))
	p("Join:	", str.Join([]string{"a","b", "c"}, "-"))
	p("Repeat:	", str.Repeat("a",5))
	p("Replace:	", str.Replace("foo", "o", "0", -1))
	p("Replace:	", str.Replace("foo", "o", "0", 1))
	p("Split:	", str.Split("a-b-c-d-e", "-"))
	p("ToLower:	", str.ToLower("TEST"))
	p("ToUpper:	", str.ToUpper("test"))
}