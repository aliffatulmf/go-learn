package main

import "fmt"

func main() {
	name := "Aliffatul"

	switch name {
	case "Alif":
		fmt.Println("Hello", name)
	case "Unknown":
		fmt.Println("Hi")
		fmt.Println(name)
	default:
		fmt.Println("Hi", name)
	}

	// short statement
	switch l := len(name); l > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
		break
		// fmt.Println("Invisible")
	case false:
		fmt.Println("Hi", name)
	}

	status := "student"
	switch {
	case status == "student":
		fmt.Println("Student")
	case status == "teacher":
		fmt.Println("Teacher")
	default:
		fmt.Println("Unknown")
	}
}
