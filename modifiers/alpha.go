package process

import "strings"

// Helper function that validate the characters starting wih a vowel or h both lowercase and uppercase
func vowel(s string) bool {
	if s == "" {
		return false
	}
	// vowels to check.
	vowel := "aeouihHAEOUI"
	return strings.ContainsRune(vowel, rune(s[0]))
}

// function that handles the vowel in your input
func Alpha(text []string) ([]string, error) {
	for i := 0; i < len(text); i++ {
		if vowel(string(text[i][0])) && i > 0 {
			if text[i-1] == "a" || text[i-1] == "A" && i > 0 {
				if text[i-1] == "a" {
					text[i-1] = "an"
				} else {
					text[i-1] = "An"
					text = append(text, text[i-1])
				}

			}
		}
	}
	return text, nil
}
