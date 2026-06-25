package practice

import (
	"fmt"
	"strconv"
)

func BinaryToDecimal(code string) (int64, error) {
	ans, err := strconv.ParseInt(code, 2, 10)
	if err != nil {
		return 0, fmt.Errorf("error")
	}
	return ans, err
}
