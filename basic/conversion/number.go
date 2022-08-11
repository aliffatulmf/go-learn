package main

import "fmt"

func main() {
	var inumber1 int8 = 100
	var inumber2 int8 = -127

	var unumber1 uint8 = uint8(inumber1)
	var unumber2 uint8 = uint8(inumber2)

	fmt.Println(inumber1, unumber1)
	fmt.Println(inumber2, unumber2)

	var ft1 float32 = 5.334299
	var ft2 float32 = 7.590490
	var res_float float64 = float64(ft1 + ft2)

	fmt.Println(res_float)
	fmt.Println(uint16(res_float))
}
