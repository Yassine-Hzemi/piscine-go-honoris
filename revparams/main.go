package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	argumentsLength := 0
	for range arguments {
		argumentsLength++
	}
	for i := argumentsLength - 1; i > 0; i-- {
		if i != 0 {
			for _, argLetter := range arguments[i] {
				z01.PrintRune(argLetter)
			}
			z01.PrintRune('\n')
		}
	}
}
