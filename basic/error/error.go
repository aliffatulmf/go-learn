package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai, pembagi uint) (uint, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	}

	res := nilai / pembagi
	return res, nil
}

func main() {
	hasil, err := Pembagian(50, 0)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Hasil :", hasil)
}
