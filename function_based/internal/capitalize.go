package practice

func Capitalize(s string) string {
	organize := make([]string, 0)
	count := 0
	tempHolder := ""
	for i, ch := range s {
		tempHolder += string(ch)
		if ch == ' ' || i == len(s)-1 {
			organize = append(organize, tempHolder)
			count++
			tempHolder = ""
			continue
		}
	}
	output := ""
	for _, ch := range organize {
		output += ToUppercase(string(ch[0])) + ToLowercase(ch[1:]) + " "
	}
	return output[:len(output)-1]
}