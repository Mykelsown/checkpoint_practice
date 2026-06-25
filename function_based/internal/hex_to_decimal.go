package practice

import (
	"fmt"
	"math"
)

// "strconv"

func HexToDecimal(code string) (int64, error) {
	// ans, err := strconv.ParseInt(code, 16, 10)
	// if err != nil {
	// 	return 0, fmt.Errorf("error")
	// }
	// return ans, err
	
	var maxVal int64 = 255
	var val int64

	for i := len(code)-1; i >= 0; i-- {
		hexVal := map[byte]int{
			'0': 0,
			'1': 1,
			'2': 2,
			'3': 3,
			'4': 4,
			'5': 5,
			'6': 6,
			'7': 7,
			'8': 8,
			'9': 9,
			'A': 10,
			'B': 11,
			'C': 12,
			'D': 13,
			'E': 14,
			'F': 15,
		}
		conv := hexVal[code[i]] * int(math.Pow(16, float64(len(code)-i-1)))
		val += int64(conv)
	}
	if val > maxVal {
		return 0, fmt.Errorf("error")
	}
	return val, nil
}
