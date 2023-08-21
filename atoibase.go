package piscine

func AtoiBase(s string, base string) int {
	baseF := Lens(base)
	result := []rune(base)
	pos := 0
	if constrains1(baseF, result) {
		for i := 0; i < Lens(s); i++ {
			pos *= Lens(base)
			pos += Index1(s[i:], base)
		}

		return pos
	}
	return 0
}

func constrains1(baseF int, result []rune) bool {
	if baseF == 1 {
		return false
	}
	for a := 0; a < Lens(string(result)); a++ {
		for x := a + 1; x < Lens(string(result)); x++ {
			if baseF < 2 || result[a] == result[x] || result[a] == '+' || result[a] == '-' {
				return false
			}
		}
	}
	return true
}

func Index1(s string, toFind string) int {
	arrayStr := []rune(s)
	array := []rune(toFind)
	for j := 0; j < Lens(s); j++ {
		for x := 0; x < Lens(toFind); x++ {
			if arrayStr[j] == array[x] {
				return x
			}
		}
	}
	return -1
}

func Lens(r interface{}) int {
	count := 0
	switch a := r.(type) {
	case []string:
		for range a {
			count++
		}
	case string:
		for range a {
			count++
		}
	}

	return count
}
