package piscine

func IsLower(str string) bool {
	if str == "" {
		return false
	}
	for _, letter := range str {
		if !(letter > 96 && letter < 123) {
			return false
		}
	}
	return true
}
