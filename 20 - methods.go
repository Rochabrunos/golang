package main

import "fmt"

type rect struct {
	width, height int
}

/*
A parte importante é a parte entre parênteses antes do nome da função: (r *rect). 
Essa parte é conhecida como o "receiver" (receptor) do método. 
No caso, (r *rect) indica que o método area é um método associado ao tipo *rect (ponteiro para rect). 
O *rect é o tipo do receiver, e r é o nome da variável que representa a instância da struct.

*/
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}