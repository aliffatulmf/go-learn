package main

import "fmt"

func Random() interface{} {
	return "Boo"
}

func main() {
	random := Random()
	randomString := random.(string)

	fmt.Println(randomString)

	// safe way
	switch random.(type) {
	case string:
		fmt.Println("random type is string")
	case int:
		fmt.Println("random type is int")
	case bool:
		fmt.Println("random type is boolean")
	default:
		fmt.Println("Unknown")
	}
}
