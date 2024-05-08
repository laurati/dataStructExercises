package main

import "fmt"

func main() {

	var p *int
	p = new(int) // Aloca memória para um int e o inicializa com zero

	fmt.Println(*p) //0
	*p = 10

	fmt.Println(*p) //10

	a := new(int)
	fmt.Println(*a)

	type Person struct {
		Name string
		Age  int
	}

	var pe *Person = new(Person)
	fmt.Println(*pe)
}
