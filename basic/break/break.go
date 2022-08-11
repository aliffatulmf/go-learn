package main

import "fmt"

func main() {
	loop := []string{"a", "b", "c", "d"}

	for i := 0; i < len(loop); i++ {
		if i == 2 {
			break
		}

		fmt.Println(i, "=", loop[i])
	}

	fmt.Println()

	for idx, val := range loop {
		if idx == 2 {
			break
		}

		fmt.Println(idx, "=", val)
	}
}
