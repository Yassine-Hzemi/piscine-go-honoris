package piscine

func IsNumeric(str string) bool {
	if StrLen(str) > 0 {
		for _, letter := range str {
			if !(letter > 47 && letter < 58) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
