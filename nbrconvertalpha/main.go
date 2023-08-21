package main

import (
	"os"

	"github.com/01-edu/z01"
)

func strToInt(s string) int {
	str := 0
	signCount := 0
	addMinus := false
	for i, v := range s {
		if int(rune(v)) > 47 && int(rune(v)) < 58 || v == '+' || v == '-' {
			if (v == '-' || v == '+') && i > 0 {
				return 0
			}
			if v == '-' && i == 0 {
				addMinus = true
			}
			if v == '+' || v == '-' {
				signCount++
				if signCount > 1 {
					return 0
				}
				continue
			}
			a := 0
			for i := '1'; i <= v; i++ {
				a++
			}
			str = str*10 + a
		} else {
			return 0
		}
	}
	if addMinus {
		str = str - (str * 2)
	}
	return str
}

func main() {
	arguments := os.Args[1:]
	i := 0
	for range arguments {
		i++
	}
	if i == 0 {
		return
	}
	isUpperFlag := false
	for _, arg := range arguments {
		if arg == "--upper" || arg == "-upper" {
			isUpperFlag = true
			arguments = os.Args[2:]
			break
		}
	}
	for _, arg := range arguments {
		number := strToInt(arg)
		if number > 0 && number < 27 {
			if isUpperFlag {
				z01.PrintRune(rune(number + 64))
			} else {
				z01.PrintRune(rune(number + 96))
			}

		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
