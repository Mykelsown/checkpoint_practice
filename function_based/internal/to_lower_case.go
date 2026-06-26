package practice

func ToLowercase(s string) string{
	res := ""
	for i:= 0; i < len(s); i++{
		if s[i] >= 'A' && s[i] <= 'Z'{
			res += string(s[i]+32)
			continue
		} 
		res += string(s[i])
	}
	return res
}