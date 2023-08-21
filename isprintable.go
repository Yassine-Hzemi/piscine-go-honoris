package piscine

func IsPrintable(str string) bool {
	if StrLen(str) > 0 {
		for _, letter := range str {
			if letter > 0 && letter < 32 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
