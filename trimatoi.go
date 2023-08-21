package piscine

func TrimAtoi(s string) int {
	numbers := 0
	isMinus := false
	blank := true
	for _, letter := range s {
		if letter == 45 && blank && !isMinus {
			isMinus = true
		} else if letter > 47 && letter < 58 {
			numbers *= 10
			numbers += int(letter - 48)
			blank = false
		}
	}
	if blank {
		return 0
	}
	if isMinus {
		return -numbers
	}
	return numbers
}
