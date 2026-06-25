package main

import (
	"fmt"
	"functionbased/internal"
)

func main() {
	fmt.Println(practice.HexToDecimal("1E"))
	fmt.Println(practice.HexToDecimal("FF"))
	fmt.Println(practice.HexToDecimal("0"))
	fmt.Println(practice.HexToDecimal("XYZ"))
}