package process

import "strconv"

func BaseConv(token []string) []string {
	var number []string

	for _, cha := range token {
		if cha == "(hex)" || cha == "(bin)" {
			base := 16
			if cha == "(bin)" {
				base = 2
			}
			if len(number) > 0 {
				num, err := strconv.ParseInt(number[len(number)-1], base, 64)
				if err == nil {
					number[len(number)-1] = strconv.FormatInt(num, 10)
				}
			}
		} else {
			number = append(number, cha)
		}
	}
	return number
}
