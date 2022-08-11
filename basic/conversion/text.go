package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi("alif")
	b, _ := strconv.Atoi("fatul")

	fmt.Println("a", a)
	fmt.Println("b", b)

	char_a := byte('a')
	char_b := byte('b')

	fmt.Println(string(char_a))
	fmt.Println(string(char_b))
}
