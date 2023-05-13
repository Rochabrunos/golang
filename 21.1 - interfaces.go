package main
import "fmt"

/* 
Source: https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
Understanding a little more about interfaces
-> A interface is two things
	-> It is a set of methods
	-> Also, it is a type
Always hungry, keep walking: https://research.swtch.com/interfaces
*/

/* 
-> Instead of designing out abstraction in terms 
of what king of data out types can hold, rather 
we design our abstraction in terms of what actions 
out types can execute
-> We define an Animal is any type that has a method 
named Speak()
-> Is said that the type satisfy the Animal interface
*/
type Animal interface {
	Speak() string
}
// All these structs satisfy the Animal interface
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

/*
Change only (c Cat -> c *Cat) this function will not work, we have to pass the *Cat pointer to our animals slice 
in the main function
*/
func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "????"
}

/* 
Let's see the empty interface, that every type satisfy the interface
The v parameter has "interface{}" type
Go runtime will perform the type conversion (if necessery),
and convert the "value" to an "interface{}" value 
*/
func DoSomething(v interface{}) {
	fmt.Printf("This interface has the %T type\n", v)
}

func PrintAllVals (v []interface{}) {
	for _, vals := range v {
		fmt.Println(vals)
	}
}

func main() {
	/*
	This works because a pointer type can access the methods of its associated value type, but not vice versa, this means that I can pass a *Dog where it is expeceted a Dog, but can't pass a Cat where *Cat is expected 
	*/
	animals := []Animal{new(Dog), new(Cat), Llama{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	DoSomething(animals)
	for _, animal := range animals {
		DoSomething(animal)
	}

	names := []string{"bruno", "lucas", "marcos"}
	/*
	This will return an error " cannot use names (variable of type []string) as []interface{} value in argument to PrintAllVals "
	PrintAllVals(names)
	In this case we hare explicitly cast our []string into []interfaces{} , here we go:
	*/
	newNames := make([]interface{}, len(names))
	for index, value := range names {
		newNames[index] = value
	}
	//Then call the PrintAllVals using "newNames" instead "names"
	PrintAllVals(newNames)
}