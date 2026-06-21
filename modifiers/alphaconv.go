package process

import (
	"strconv"
	"strings"
)

// function that convert strings to lower or uppercase and convert a first letter of a string to uppercase.
func AlphaConv(token []string) ([]string, error) {
	var new_aplha []string
	for _, word := range token {
		if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") {
			char := word[1 : len(word)-1]
			part := strings.Split(char, ",")
			number := 1
			part1 := strings.TrimSpace(part[0])
			// validating when the content in the bracket are more that one
			if len(part) > 1 {
				num, err := strconv.Atoi(strings.TrimSpace(part[1]))
				if err == nil {
					if num > number {
						number = num
					}
				}
			}
			start := len(new_aplha) - number
			if start < 0 {
				start = 0
			}
			for i := start; i < len(new_aplha); i++ {
				new_part1 := strings.ToLower(part1)
				switch new_part1 {
				case "up":
					new_aplha[i] = strings.ToUpper(new_aplha[i])
				case "low":
					new_aplha[i] = strings.ToLower(new_aplha[i])
				case "cap":
					news := []rune(new_aplha[i])
					if news[0] >= 'a' && news[0] <= 'z' {
						news[0] = news[0] - 32
					}
					new_aplha[i] = string(news)
				}
			}

		} else {
			new_aplha = append(new_aplha, word)
		}
	}
	return new_aplha, nil
}
