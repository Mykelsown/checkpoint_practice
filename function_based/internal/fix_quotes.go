package practice

func FixQuotes(str string) string {
	isOpen := map[string]bool{"'": true}
	words := toStrSlice(str)
	res := ""

	for _, word := range words {
		if isOpen[word] {
			res += word
			isOpen[word] = false
			continue
		}
		
		if !isOpen[word] && word == "'" {
			res = res[:len(res)-1]
			isOpen[word] = true
		}

		res += word + " "
	}

	return res
}

func toStrSlice(s string) []string {
	res := make([]string, 0)
	elementTempBuild := ""
	for _, ch := range s + " " {
		if ch == ' ' {
			res = append(res, elementTempBuild)
			elementTempBuild = ""
			continue
		}
		elementTempBuild += string(ch)
	}
	return res
}

func joinStrSlice(words []string) string {
	res := ""
	for _, word := range words {
		res += word + " "
	}
	return res[:len(res)-1]
}
