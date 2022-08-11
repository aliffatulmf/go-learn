package main

import "fmt"

func main() {
	function_variadic("Variadic", "alpha", "beta", "charlie", "delta")

	slice := []string{
		"alpha", "beta", "charlie", "delta",
	}
	function_variadic("Variadic 2", slice...)
}

func function_variadic(a string, b ...string) {
	fmt.Println("Header", a)

	for i := 0; i < len(b); i++ {
		fmt.Println("Line", i, "=", b[i])
	}
}

// ❌ bad example
// func bad_example(a ...string, b string) {
// 	fmt.Println(a, b)
// }
