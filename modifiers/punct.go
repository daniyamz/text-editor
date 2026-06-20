package process

import "strings"

// helper function that validate punctuations
func puncta(p string) bool {
	if p == "" {
		return false
	}
	punc := "!?.,:;"
	return strings.ContainsRune(punc, rune(p[0]))
}

// Function that handle puntuation marks in the inpurt
func PunctControl(str []string) string {
	for x := 0; x < len(str); x++ {
		if puncta(str[x]) && x > 0 {
			str[x-1] = str[x-1] + str[x]
			str = append(str[:x], str[x+1:]...)
		}
	}
	return strings.Join(str, " ")
}

// function that handle the quotetion in the input.
func QuotControl(txt string) (string, error) {
	text := strings.Split(txt, "'")
	for t := 1; t < len(text); t += 2 {
		text[t] = strings.TrimSpace(text[t])
	}
	return strings.Join(text, "'"), nil
}
