package piscine

func ToLower(s string) string {
	str := []rune(s)
	for i, letter := range s {
		if letter > 64 && letter < 91 {
			str[i] = letter + 32
		} else {
			str[i] = letter
		}
	}
	return string(str)
}
