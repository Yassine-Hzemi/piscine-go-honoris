package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	isLetter := false
	firstUpper := false
	i := 0
	for range runes {
		if (runes[i] >= 65 && runes[i] <= 90) || (runes[i] >= 97 && runes[i] <= 122) || (runes[i] >= 48 && runes[i] <= 57) {
			if isLetter == false {
				firstUpper = true
				isLetter = true
			} else {
				firstUpper = false
			}
		} else {
			firstUpper = false
			isLetter = false
		}
		if isLetter == true {
			if firstUpper == false {
				if runes[i] >= 65 && runes[i] <= 90 {
					runes[i] += 32
				}
			} else {
				if runes[i] >= 97 && runes[i] <= 122 {
					runes[i] -= 32
				}
			}
		}
		i++
	}
	return string(runes)
}
