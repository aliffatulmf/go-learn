package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "Semarang",
		Province: "Jawa Tengah",
		Country:  "Indonesia",
	}
	address2 := &address1
	address2.City = "Wonogiri"

	*address2 = Address{
		City:     "Surakarta",
		Province: "Jawa Tengah",
		Country:  "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)

	var newAddr *Address = new(Address)
	newAddr.City = "Salatiga"

	fmt.Println(newAddr)
}
