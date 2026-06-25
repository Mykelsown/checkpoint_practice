package practice

import (
	"fmt"
	// "strconv"
	"math"
)

func BinaryToDecimal(code string) (int64, error) {
	// // Solved with strconv package
	// ans, err := strconv.ParseInt(code, 2, 10)
	// if err != nil {
	// 	return 0, fmt.Errorf("error")
	// }
	// return ans, err

	// without strconv package
	var conv int64
	for i := len(code)-1; i >= 0; i-- {
		if code[i] != '0' && code[i] != '1' {
			return 0, fmt.Errorf("error")
		}
		// conv is the whole conversion for each individual charcter in the string to the desired format
		// the number 48 there is the difference between the byte to the actual integer value for a valid binary i.e byte(0) to decimal -> 48 - 48 
		conv += int64(code[i]-48) * int64(math.Pow(2, float64(len(code)-i-1))) 
	}
	return conv, nil
}
