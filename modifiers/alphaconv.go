package process

import "strings"

func AlphaConv(token []string) []string {
	var new_aplha []string
	for _, word := range token {
		if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") {
			char := word[1 : len(word)-1]
			//char = strings.ReplaceAll(char, " ", "")
			switch char {
			case "up":
				if len(new_aplha) > 0 {
					new_aplha[len(new_aplha)-1] = strings.ToUpper(new_aplha[len(new_aplha)-1])
				}
			case "low":
				if len(new_aplha) > 0 {
					new_aplha[len(new_aplha)-1] = strings.ToLower(new_aplha[len(new_aplha)-1])
				}
			case "cap":
				if len(new_aplha) > 0 {
					var runes []rune = []rune(new_aplha[len(new_aplha)-1])
					if runes[0] >= 'a' && runes[0] <= 'z' {
						runes[0] = runes[0] - 32
						new_aplha[len(new_aplha)-1] = string(runes)
					}
				}
			}
		} else {
			new_aplha = append(new_aplha, word)
		}
	}
	return new_aplha
}
