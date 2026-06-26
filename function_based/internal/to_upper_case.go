package practice

func ToUppercase(s string) string{
	res := ""
	for i:= 0; i < len(s); i++{
		if s[i] >= 'a' && s[i] <= 'z'{
			res += string(s[i]-32)
			continue
		} 
		res += string(s[i])
	}
	return res
}