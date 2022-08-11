package main

import "fmt"

func main() {
	name := "Aliffatul"
	defer stopApps()

	if name != "Aliffatul" {
		fmt.Println("Not Aliffatul")
	} else {
		fmt.Println("Hi,", name)
	}
}

func stopApps() {
	fmt.Println("stopped")
}
