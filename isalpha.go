package piscine

func IsAlpha(str string) bool {
	if StrLen(str) > 0 {
		for _, letter := range str {
			if !(letter > 47 && letter < 58) && !(letter > 64 && letter < 91) && !(letter > 96 && letter < 123) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
