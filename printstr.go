package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for index := range str {
		z01.PrintRune(rune(str[index]))

	}
}
