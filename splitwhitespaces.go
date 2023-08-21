package piscine

func SplitWhiteSpaces(str string) []string {
	var len = 0
	var strlLen = 0
	var start = 0
	var j = 0
	var wasLetter = false
	for _, v := range str {
		if (v == '\n' || v == '\t' || v == ' ') && wasLetter {
			len++
			wasLetter = false
		} else if v != '\n' && v != '\t' && v != ' ' {
			wasLetter = true
		}
		strlLen++
	}
	if wasLetter {
		len++
		wasLetter = false
	}
	arr := make([]string, len)
	for i, v := range str {
		if (v == '\n' || v == '\t' || v == ' ') && wasLetter {
			arr[j] = str[start:i]
			j++
			wasLetter = false
		} else if v != '\n' && v != '\t' && v != ' ' {
			if !wasLetter {
				start = i
			}
			wasLetter = true
		}
	}
	if wasLetter {
		arr[j] = str[start:strlLen]
	}
	return arr
}
