package piscine

import (
	"github.com/01-edu/z01"
)

// PrintNbrBase  prints an int in a string base passed in parameters.
func PrintNbrBase(nbr int, base string) {
	if nbr == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
		return
	}
	len := 0

	sign := '+'
	if nbr < 0 {
		sign = '-'
		nbr *= -1
	}

	for range base {
		len++
	}

	// fmt.Println(pow)

	// if nbr < len*len {
	// 	z01.PrintRune('N')
	// 	z01.PrintRune('V')
	// 	return
	// }

	if len < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {

			if base[i] == base[j] && i != j {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			} else if base[i] == '+' || base[i] == '-' {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	str := toBase(nbr, len)

	arr := []byte(str)

	for i := range arr {
		for x := 0; x < len; x++ {
			if int(arr[i]-48) == x {
				arr[i] = base[x]
			}
		}
	}

	if sign == '-' {
		z01.PrintRune('-')
	}
	for _, c := range arr {
		z01.PrintRune(rune(c))
	}

}

func toBase(nb, base int) string {

	based := ""
	for nb != 0 {
		based += string(nb%base + 48)
		nb = nb / base
	}

	len := 0
	for range based {
		len++
	}

	reversed := []rune(based)
	for i, c := range based {
		reversed[len-1-i] = c
	}

	return string(reversed)
}
