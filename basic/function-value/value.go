package main

import "fmt"

func main() {
	name := setName

	sayHi(name("Alif"))
}

func setName(n string) string {
	return fmt.Sprint("my name is ", n)
}

func sayHi(n string) {
	fmt.Println("Hi", n)
}
