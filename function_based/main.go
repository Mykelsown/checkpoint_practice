package main

import (
	"fmt"
	practice "functionbased/internal"
)

func main() {
	fmt.Println(practice.HexToDecimal("1E"))
	fmt.Println(practice.HexToDecimal("FF"))
	fmt.Println(practice.HexToDecimal("0"))
	fmt.Println(practice.HexToDecimal("XYZ"))
	// fmt.Println(practice.BinaryToDecimal("1010"))
	// fmt.Println(practice.BinaryToDecimal("1"))
	// fmt.Println(practice.BinaryToDecimal("0"))
	// fmt.Println(practice.BinaryToDecimal("102"))
}
