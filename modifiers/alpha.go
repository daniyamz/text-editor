package process

import "strings"

func vowel(s string) bool {
	if s == "" {
		return false
	}
	vowel := "aeouihHAEOUI"
	return strings.ContainsRune(vowel, rune(s[0]))
}

func Alpha(text []string) []string {
	for i := 0; i < len(text); i++ {
		if vowel(string(text[i][0])) && i > 0 {
			if text[i] == "a" || text[i] == "A" && i > 0 {
				if text[i] == "a" {
					text[i] = "an"
				} else {
					text[i] = "An"
					text = append(text, text[i])
				}

			}
		}
	}
	return text
}
