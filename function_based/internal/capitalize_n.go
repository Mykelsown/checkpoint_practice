package practice

import "fmt"

func CapitalizeN(words []string, n int) []string {
	if n > len(words){
		n = 0
	}
	for i := len(words) - 1; i >= 0; i-- {
		upperCase := ""
		if words[i][0] >= 'a' && words[i][0] <= 'z' && i >=n {
			upperCase = string(words[i][0]-32) + words[i][1:]
			words[i] = fmt.Sprintf("%q", upperCase)
			continue
		}
		words[i] = fmt.Sprintf("%q", ToLowercase(words[i]))
	}

	return words
}