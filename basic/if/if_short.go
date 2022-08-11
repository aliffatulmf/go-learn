package main

import "fmt"

func main() {
	if name := "Fatul"; len(name) > 5 {
		fmt.Println("Nama lebih dari 5 karakter")
	} else {
		fmt.Println("Nama saya", name)
	}
}
