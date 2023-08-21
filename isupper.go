package piscine

func IsUpper(str string) bool {
	if StrLen(str) > 0 {
		for _, letter := range str {
			if !(letter > 64 && letter < 91) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
