package piscine

func Atoi(str string) int {
	num := 0
	znak := 1

	for i, val := range str {
		digit := int(val)
		digit -= 48

		if digit <= 9 && digit >= 0 {
			num = num*10 + digit
		} else if digit == -3 && i == 0 {
			znak = -1
		} else if digit == -5 && i == 0 {
			znak = 1
		} else {
			return 0
		}
	}

	num *= znak

	return num
}
