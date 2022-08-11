package main

import "fmt"

func main() {
	counter := 0

	change := func(n int) {
		counter = n
	}

	change(10)

	fmt.Println(counter)
}
