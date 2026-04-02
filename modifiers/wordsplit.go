package process

func ispunt(s string) bool {
	switch s {
	case ",", ":", ";", "?", ".", "!":
		return true
	}
	return false
}

func Split(str string) []string {
	var word string
	var token []string
	for _, char := range str {
		if char == ' ' {
			if word != "" {
				token = append(token, word)
				word = ""
			}
		} else if ispunt(string(char)) {
			if word != "" {
				token = append(token, word)
				word = ""
			}
			token = append(token, string(char))
		} else {
			word += string(char)
		}
	}
	if word != "" {
		token = append(token, word)
	}
	return token
}
