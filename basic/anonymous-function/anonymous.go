package main

import "fmt"

type Blacklist func(name string) bool

func main() {
	registerUser("alif", func(name string) bool {
		if name == "admin" {
			return true
		}

		return false
	})
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println(name, "blocked")
	} else {
		fmt.Println("Welcome,", name)
	}
}
