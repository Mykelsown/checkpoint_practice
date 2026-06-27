package practice

func ReverseString(s string) string{
	res := ""
	for _, ch := range s {
		res = string(ch) + res
	}
	return res
}