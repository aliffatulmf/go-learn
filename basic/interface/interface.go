package main

import "fmt"

type HasName interface {
	GetName() string
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	cat := Animal{
		Name: "Meong",
	}

	SayHello(cat)
}

func SayHello(hasName HasName) {
	fmt.Println("Hello,", hasName.GetName())
}
