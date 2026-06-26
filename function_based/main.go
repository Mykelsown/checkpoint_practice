package main

import (
	"fmt"
	practice "functionbased/internal"
)

func main() {
	// // SECTION 1: NUMBERS
	// println("===========Hex============")
	// fmt.Println(practice.HexToDecimal("1E"))
	// fmt.Println(practice.HexToDecimal("FF"))
	// fmt.Println(practice.HexToDecimal("0"))
	// fmt.Println(practice.HexToDecimal("XYZ"))
	// println("===========Bin============")
	// fmt.Println(practice.BinaryToDecimal("1010"))
	// fmt.Println(practice.BinaryToDecimal("1"))
	// fmt.Println(practice.BinaryToDecimal("0"))
	// fmt.Println(practice.BinaryToDecimal("102"))
	// println("===========Oct============")
	// fmt.Println(practice.OctalToDecimal("17"))
	// fmt.Println(practice.OctalToDecimal("100"))
	// fmt.Println(practice.OctalToDecimal("89"))
	// println("===========Dec============")
	// fmt.Println(practice.DecimalToBase(30, 16))
	// fmt.Println(practice.DecimalToBase(10, 2))
	// fmt.Println(practice.DecimalToBase(0.3, 2))
	// fmt.Println(practice.DecimalToBase(1341, 16))
	// fmt.Println(practice.DecimalToBase(15, 10))

	// // SECTION 2: String Maniplation
	println("===========ToUpperCase============")
	fmt.Println(practice.ToUppercase("hello"))
	fmt.Println(practice.ToUppercase("Go Lang"))
	fmt.Println(practice.ToUppercase("123!abc"))
}
