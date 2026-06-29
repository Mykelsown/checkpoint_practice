package practice

func IsPalindrome(s string) bool {
	s = ToLowercase(s)
	revSecondHalf := ""
	for i := len(s)-1; i > len(s)/2; i-- {
		revSecondHalf += string(s[i])
	} 
	return s[0:len(s)/2] == revSecondHalf
}