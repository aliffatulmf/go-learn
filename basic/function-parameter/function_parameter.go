package main

import "fmt"

func main() {
	print_text("Alif")
	loop_start_to_end(1, 10)
	loop_start_to_end(100, 10)
}

func print_text(t string) {
	fmt.Println("print_text:", t)
}

func loop_start_to_end(start, end int) {
	if start > end {
		fmt.Println("start greater than end")
	} else {
		for i := start; i < end; i++ {
			fmt.Println(i)
		}
	}
}
