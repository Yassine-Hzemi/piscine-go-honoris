package main

import (
	"github.com/01-edu/z01"
)

func main() {

	for i := 'a'; i <= 'z'; i++ {

		z01.PrintRune(rune(i))

	}
	z01.PrintRune(rune("\n"))
}