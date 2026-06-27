package practice

func FixPunctuation(s string) string {
	cont := make([]string, 0)
	punctuations := map[rune]bool{'.': true, ',': true, '!': true, '?': true, ':': true, ';': true}
	for i, ch := range s {
		// check for current character being a punctuation and the next a space
		if punctuations[ch] && i < len(s)-1 && i > 0 {
			if !(s[i+1] == ' ') {
				cont = append(cont, string(ch)+" ")
				continue
			}
		}

		// check for current character being a space and then removes every trailing space, including  the current character
		if ch == ' ' && i < len(s)-1 && i > 0 {
			if punctuations[rune(s[i+1])] {
				continue
			} else if s[i+1] == ' ' {
				continue
			} 
		}
		cont = append(cont, string(ch))
	}

	output := ""
	// Convert from slice to string
	for _, ch := range cont {
		output += string(ch)
	}

	return output
}