package main

import "fmt"

func main() {
	number := 0

	if number == 0 {
		panic("Number error")
	} else {
		fmt.Println("Success")
	}
}
