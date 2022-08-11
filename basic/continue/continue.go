package main

import "fmt"

func main() {
	loop := []string{"a", "b", "c", "d"}

	for i := 0; i < len(loop); i++ {
		if i == 2 {
			continue
		}

		fmt.Println(i, "=", loop[i])
	}

	fmt.Println()

	for idx, val := range loop {
		if idx == 2 {
			continue
		}

		fmt.Println(idx, "=", val)
	}
}
