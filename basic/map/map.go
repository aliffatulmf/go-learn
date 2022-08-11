package main

import "fmt"

func main() {
	maps := map[string]int{
		"Satu": 1,
	}
	maps["Dua"] = 2
	maps["Tiga"] = 3
	maps["Lima"] = 5

	fmt.Println(maps)

	maps["Tiga"] = 33
	fmt.Println(maps)

	delete(maps, "Lima")
	fmt.Println(maps)
}
