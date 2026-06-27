package process

import (
	"strings"
)

// Helper function that render punctations in an input.
func Ispuct(s rune) bool {

	punt := ",./?':![]"
	return strings.ContainsRune(punt, s)
}

// function that convert input strings into a slice of string make each character a token.
// It handles the punctuations from the helper function above
// It handles bracket too making them a standalone token
func Split(str string) ([]string, error) {
	var token []string
	//var word string
	var words strings.Builder

	for _, char := range str {
		switch {
		case char == '(':
			if words.Len() != 0 {
				token = append(token, words.String())
				words.Reset()
			}
			token = append(token, string(char))
			words.Reset()
		case char == ')':
			if words.Len() != 0 {
				token = append(token, words.String())
				words.Reset()
			}
			token = append(token, string(char))
			words.Reset()
		case Ispuct(char):
			if words.Len() != 0 {
				token = append(token, words.String())
				words.Reset()
			}
			token = append(token, string(char))
		case char == ' ':
			if words.Len() != 0 {
				token = append(token, words.String())
				words.Reset()
			}
		default:

			words.WriteString(string(char))
		}
	}
	if words.Len() != 0 {
		token = append(token, words.String())
		words.Reset()
	}

	return token, nil
}
