package main

import (
	"fmt"
	practice "functionbased/internal"
)

func main() {
	println("===========Hex============")
	fmt.Println(practice.HexToDecimal("1E"))
	fmt.Println(practice.HexToDecimal("FF"))
	fmt.Println(practice.HexToDecimal("0"))
	fmt.Println(practice.HexToDecimal("XYZ"))
	println("===========Bin============")
	fmt.Println(practice.BinaryToDecimal("1010"))
	fmt.Println(practice.BinaryToDecimal("1"))
	fmt.Println(practice.BinaryToDecimal("0"))
	fmt.Println(practice.BinaryToDecimal("102"))
	println("===========Oct============")
	fmt.Println(practice.OctalToDecimal("17") )  
	fmt.Println(practice.OctalToDecimal("100"))  
	fmt.Println(practice.OctalToDecimal("89") )
}
