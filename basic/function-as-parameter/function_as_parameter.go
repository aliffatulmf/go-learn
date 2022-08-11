package main

import (
	"fmt"
)

func main() {
	print_name(print_text, "Aliffatul")

	bad := []string{"anjing", "manuk", "kambing", "tupai"}
	text := []string{"aku", "suka", "kambing", "dan", "anjing", "mungkin", "tupai"}

	process(text_filtering, text, bad)

}

func print_text(n string) {
	fmt.Println("My name is", n)
}

func print_name(fn func(n string), name string) {
	fn(name)
}

func text_filtering(text []string, bad []string) map[string][]string {
	var gd []string
	var bd []string

	filtered := make(map[string][]string, 2)

	for _, t := range text {
		if contain(t, bad) {
			bd = append(bd, t)
		} else {
			gd = append(gd, t)
		}
	}

	filtered["good"] = gd
	filtered["bad"] = bd

	return filtered
}

func contain(text string, bad []string) bool {
	for _, b := range bad {
		if text == b {
			return true
		}
	}

	return false
}

func process(tf func(text []string, bad []string) map[string][]string, text, bad []string) {
	result := tf(text, bad)

	fmt.Println("Good", result["good"])
	fmt.Println("Bad", result["bad"])
}
