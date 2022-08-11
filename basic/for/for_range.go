package main

import "fmt"

func main() {
	var slice []int

	for i := 0; i < 50; i++ {
		slice = append(slice, i)
	}

	fmt.Println(slice)

	for index, value := range slice {
		fmt.Println(index, value)
	}
}
