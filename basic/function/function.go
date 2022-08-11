package main

import "fmt"

func main() {
	print_hello()
	print_world()
	do_loop_1_to_10()
	do_print_hello_func_10_times()
}

func print_hello() {
	fmt.Println("Hello")
}

func print_world() {
	fmt.Println("World")
}

func do_loop_1_to_10() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func do_print_hello_func_10_times() {
	for i := 1; i <= 10; i++ {
		print_hello()
	}
}
