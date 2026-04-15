package process

import "strings"

func Ispuct(s string) bool {

	punt := ",./?':![]"
	return strings.ContainsRune(punt, rune(s[0]))
}

func Split(str string) []string {
	var token []string
	var word string
	count := 0
	for _, char := range str {

		if char == '(' {
			if word != "" && count == 0 {
				token = append(token, word)
				word = ""
			}
			count++
			word += string(char)
		} else if char == ')' {
			count--
			word += string(char)
			if count == 0 && word != "" {
				token = append(token, word)
				word = ""
			}
		} else if Ispuct(string(char)) && count == 0 {
			if word != "" && !Ispuct(string(word[len(word)-1])) {
				token = append(token, word)
				word = ""
			}
			word += string(char)
		} else if char == ' ' && count == 0 {
			if word != "" {
				token = append(token, word)
				word = ""
			}
		} else {
			if word != "" && Ispuct(string(word[len(word)-1])) && count == 0 {
				token = append(token, word)
				word = ""
			}
			word += string(char)
		}
	}
	if word != "" {
		token = append(token, word)
	}
	return token
}
