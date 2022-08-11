package main

import (
	"fmt"
)

func main() {
	input := 0

	defer func() {
		message := recover()
		if message != nil {
			fmt.Println("Error message:\t", message)
		}
	}()

	result := 10 / input

	fmt.Println(result)
}
