package main

import "fmt"

func main() {
	fname, lname := named_return_value()

	fmt.Println(fname, lname)
}

func named_return_value() (fname, lname string) {
	fname = "Alif"
	lname = "Fatul"

	return
}
