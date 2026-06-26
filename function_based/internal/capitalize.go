package practice

func Capitalize(s string) []string{
	organize := make([]string, 0)
	count := 0
	tempHolder := ""
	for _, ch := range s{
		tempHolder += string(ch)
		if ch == ' '{
			organize = append(organize, tempHolder)
			count++
			tempHolder = ""
			continue
		}
	}
	return organize
}