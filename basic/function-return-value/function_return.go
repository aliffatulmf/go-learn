package main

import "fmt"

func main() {
	f1 := return_zero()
	f2 := return_text("Alif")

	fmt.Println(f1)
	fmt.Println(f2)

	loop_start_to_end(10, 1)
}

func return_zero() int {
	return 0
}

func return_text(t string) string {
	return "func return_text: " + t
}

func loop_start_to_end(start, end int) {
	if start > end {
		fmt.Println("start greater than end")
		return // similar to break
	}
	for i := start; i < end; i++ {
		fmt.Println(i)
	}
}
