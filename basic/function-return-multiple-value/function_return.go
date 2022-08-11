package main

import "fmt"

func main() {
	fname, lname := return_name()

	fmt.Println(fname, lname)

	first, middle, last := triple_return()

	fmt.Println(first, middle, last)
}

func return_name() (string, string) {
	return "alif", "fatul"
}

func triple_return() (string, string, string) {
	return "Aliffatul", "M", "F"
}

func return_pre_dec() (fname, lname string) {
	fname = "Alif"
	lname = "Fatul"

	return
}
