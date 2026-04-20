package process

import "strings"

func puncta(p string) bool {
	if p == "" {
		return false
	}
	punc := "!?.,:;"
	return strings.ContainsRune(punc, rune(p[0]))
}

func PunctControl(str []string) string {
	for x := 0; x < len(str); x++ {
		if puncta(str[x]) && x > 0 {
			str[x-1] = str[x-1] + str[x]
			str = append(str[:x], str[x+1:]...)
		}
	}
	return strings.Join(str, " ")
}
