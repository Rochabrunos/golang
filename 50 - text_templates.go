package main

import (
	"os"
	"text/template"
)
/*
Templates are mix of static text and actions 
enclosed in {{...}} that are used to dynamically 
insert content. Go offers built-in support for:
- createing dynamic contet
- showing customized output to the user
*/
func main() {
	t1 := template.New("t1")

	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}
	//allow error handling at compilation time
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	//We can use field names in the case of structured data
	t2 := Create("t2", "Name: {{.Name}}\n")
	t2.Execute(os.Stdout, struct{
		Name string
	}{"Jane Doe"})
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})
	//if/else provide conditional execution for templates false (0, empty string, nil pointer)
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")

	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	//loops through silcers, arrays, maps or channels
	t4 := Create("t4", "Range: {{range .}} {{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})
}