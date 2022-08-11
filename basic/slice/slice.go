package main

import "fmt"

func main() {
	// slice := []string{"Ayam", "Bebek", "Cacing"}
	slice := make([]string, 3, 5)
	slice[0] = "Ayam"
	slice[1] = "Bebek"
	slice[2] = "Cacing"
	fmt.Println("Slice\t->", slice, "Length\t->", len(slice), "Capacity\t->", cap(slice))

	fmt.Println()

	slice1 := slice[1:]
	fmt.Println("Slice 1\t->", slice1, "Length\t->", len(slice1), "Capacity\t->", cap(slice1))

	slice1[0] = "Burung" // replace from "Bebek" to "Burung"
	fmt.Println("Slice 1\t->", slice1, "Length\t->", len(slice1), "Capacity\t->", cap(slice1))

	slice = append(slice, "Domba")
	slice = append(slice, "Elang")
	fmt.Println("Slice\t->", slice, "Length\t->", len(slice), "Capacity\t->", cap(slice))

	fmt.Println()

	slice11 := slice1[1:]
	slice11 = append(slice11, "Duck", "Eagle")
	slice11[0] = "Curut"
	fmt.Println("Slice 11->", slice11, "Length\t->", len(slice11), "Capacity\t->", cap(slice11))
	fmt.Println("Slice 1\t->", slice1, "Length\t->", len(slice1), "Capacity\t->", cap(slice1))

	slice2 := make([]string, len(slice11))
	copy(slice2, slice11)

	fmt.Println("Slice 2\t->", slice2)
}
