package process

import "strings"

// Helper function that render punctations in an input.
func Ispuct(s string) bool {

	punt := ",./?':![]"
	return strings.ContainsRune(punt, rune(s[0]))
}

// function that convert input strings into a slice of string make each character a token.
func Split(str string) ([]string, error) {
	var token []string
	var word string
	count := 0
	for _, char := range str {
		// handling "(" as a single token.
		if char == '(' {
			if word != "" && count == 0 {
				token = append(token, word)
				word = ""
			}
			count++
			word += string(char)
			// handling ")" as a single token
		} else if char == ')' {
			count--
			word += string(char)
			if count == 0 && word != "" {
				token = append(token, word)
				word = ""
			}
			//handling the punctuation that as a single token.
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
	return token, nil
}
