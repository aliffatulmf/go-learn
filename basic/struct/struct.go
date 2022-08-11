package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int
}

// func Greeting(person Person, name string) {
// 	fmt.Println("Hello", person.Name, "my name is", name)
// }

func (person Person) Greeting(name string) {
	fmt.Println("Hello", person.Name, "my name is", name)
}

func main() {
	var person Person
	person.Name = "Alif"
	person.Address = "Indonesia"
	person.Age = 21

	fmt.Println(person)

	// struct literals
	john := Person{
		Name:    "John",
		Address: "Singapore",
		Age:     30,
	}
	fmt.Println(john)

	ruly := Person{"Ruly", "Indonesia", 25}
	fmt.Println(ruly)

	person.Greeting("alif")
}
