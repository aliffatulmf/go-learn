package main

import "fmt"

func main() {
	dict := map[string]int{
		"satu":  1,
		"dua":   2,
		"tiga":  3,
		"empat": 4,
	}

	for key, value := range dict {
		fmt.Println(key, "=", value)
	}
}
