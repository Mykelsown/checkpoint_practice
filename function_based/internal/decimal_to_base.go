package practice

import (
	"fmt"
)

var val = map[int]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "A",
	11: "B",
	12: "C",
	13: "D",
	14: "E",
	15: "F",
}

func DecimalToBase(n any, base int) (string, error) {
	// Base number validity check
	switch base {
	case 2, 8, 16:
		// Valid base
	default:
		return "", fmt.Errorf("error (invalid base)")
	}

	finRes := ""

	// Type Switch: Decimal validity check and conversion
	switch i := n.(type) {
	case float64:
		precision := 10
		finRes = "0."+fractionConverter(i, base, precision)

	case int:
		finRes = integerConverter(i, base)
	default:
		return "", fmt.Errorf("error (invalid decimal)")
	}
	return fmt.Sprintf("%q, ",finRes), nil
}

func integerConverter(n, base int) string {
	if n < base{
		return val[n]
	}
	
	return integerConverter(n/base, base) + val[n%base]
}
// // My own functon; but it was a little flawed
// func fractionConverter(n float64, base int) string {
// 	mult := n*float64(base)
// 	newN := mult-float64(int(mult))
// 	if n == newN {
// 		return val[int(mult)]
// 	}
// 	return val[int(mult)]+fractionConverter(newN, base)
// }

// AI modification
func fractionConverter(n float64, base int, precision int) string {
	if n < 0 || n >= 1 {
		return ""
	}

	if precision == 0 {
		return ""
	}

	mult := n * float64(base)
	digit := int(mult)

	newN := mult - float64(digit)

	return val[digit] + fractionConverter(newN, base, precision-1)
}