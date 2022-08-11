package main

import "fmt"

func main() {
	var arr [5]int = [5]int{
		1,
		2,
		3,
	}

	arr[3] = 4
	arr[4] = 5
	// ...
	// arr[n] = x -> error

	fmt.Println(arr, len(arr), cap(arr))

	arr[0] = 1_000
	fmt.Println(arr, len(arr), cap(arr))

	var arr1 = make([]int, 3)
	// arr1[0] = 1
	// arr1[1] = 2
	arr1[2] = 3

	fmt.Println(arr1, len(arr1), cap(arr1))
}
