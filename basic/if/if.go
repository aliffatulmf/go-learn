package main

import "fmt"

func main() {
	name := "Alif"

	if name == "Alif" {
		fmt.Println("Benar")
	} else {
		fmt.Println("Salah")
	}

	if name == "Fatul" {
		fmt.Println("Benar")
	} else {
		fmt.Println("Salah")
	}

	if name == "Alif" {
		fmt.Println("Nama saya alif")
	} else if name == "Fatul" {
		fmt.Println("Nama saya alif")
	} else {
		fmt.Println("Nama tidak diketahui")
	}
}
