package main

import "fmt"

func main() {
	fmt.Println(factorial(4))

	fmt.Println(recursive(1))
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func recursive(x uint) uint {
	if x > 9 {
		return 10
	}

	fmt.Print(x)

	return recursive(x + 1)
}
