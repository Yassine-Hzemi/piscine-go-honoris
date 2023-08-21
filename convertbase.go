package piscine

func IndexC(s string, c rune) int {
	str := []rune(s)
	index := -1
	for n := range str {
		if str[n] == c {
			index = n
			return index
		}
	}
	return index
}

func CheckBase(s string) bool {
	if Length(s) < 2 {
		return false
	}
	if (IndexC(s, '-') != -1) || (IndexC(s, '+') != -1) {
		return false
	}
	count := 0
	for _, a := range s {
		count++
		if IndexC(s[count:], a) >= 0 {
			return false
		}
	}
	return true
}

func Length(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func IntLen(str []int) int {
	s := 0
	for range str {
		s++
	}
	return s
}

func MyAppend(a, b []int) []int {
	p := [200]int{-1}
	count := 0
	for range a {
		p[count] = a[count]
		count++
	}
	len2 := IntLen(a)
	for range b {
		p[count] = b[count-len2]
		count++
	}
	return p[:count]
}

func ConvertBaseTo(nb int, base int, sequence []int) []int {
	if nb < 0 {
		return ConvertBaseTo(nb/-1, base, sequence)
	} else if nb > base {
		r := nb % base
		sequence = MyAppend([]int{r}, sequence)
		return ConvertBaseTo(nb/base, base, sequence)
	} else {
		r := nb % base
		sequence = MyAppend([]int{r}, sequence)
		sequence = MyAppend([]int{nb / base}, sequence)

		return sequence

	}
}

func nbrBase(nb int, base string) string {
	lenBase := Length(base)
	sequence := []int{}
	overflow := 0
	sign := ""
	if CheckBase(base) == false {
		return "NV"
	}
	if nb < 0 {
		sign = "-"
	}
	if nb > 9223372036854775807 || nb < -9223372036854775807 {
		nb = nb - 1
		overflow = 1
	}
	p := ConvertBaseTo(nb, lenBase, sequence)
	base2 := []rune(base)
	if p[0] == 0 {
		p = p[1:]
	}
	pl := IntLen(p)
	p[pl-1] = p[pl-1] + overflow
	answer := ""
	for _, n := range p {
		answer = answer + string(base2[n])
	}
	return sign + answer
}

func toInt3(a rune) int {
	s := 0
	for i := '0'; i < a; i++ {
		s++
	}
	return s
}

func StrLen3(str string) int {
	s := 0
	for range str {
		s++
	}
	return s
}

func toPower(nb int, power int) int {
	if power < 0 {
		nb = 0
	} else if power == 0 {
		nb = 1
	} else {
		nb = nb * toPower(nb, power-1)
	}
	return nb
}

func Check3(s string) bool {
	Che := true
	num := 0
	for _, n := range s {
		if (n < '0') || (n > '9') {
			num++
		}
	}

	if num > 0 {
		Che = false
	}
	return Che
}
func IntLen1(str []int) int {
	s := 0
	for range str {
		s++
	}
	return s
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	intnbr := AtoiBase(nbr, baseFrom)
	s := nbrBase(intnbr, baseTo)
	return s
}
