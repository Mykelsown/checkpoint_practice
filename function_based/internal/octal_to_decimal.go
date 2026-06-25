package practice

import (
	"fmt"
	"math"
)

func OctalToDecimal(oct string) (int, error) {
	var conv int

	for i := len(oct) - 1; i >= 0; i-- {
		octVal := map[byte]int{
			'0': 0,
			'1': 1,
			'2': 2,
			'3': 3,
			'4': 4,
			'5': 5,
			'6': 6,
			'7': 7,
		}
		if oct[i] < 0 && oct[i] > 0 {
			return 0, fmt.Errorf("error")
		}
		conv = octVal[oct[i]] * int(math.Pow(16, float64(len(oct)-i-1)))
	}
	return conv, nil
}