package practice

import (
	"fmt"
	"strconv"
)

func HexToDecimal(code string) (int64, error){
	ans, err := strconv.ParseInt(code, 16, 10)
	if err != nil {
		return 0, fmt.Errorf("error")
	}
	return ans, err
}