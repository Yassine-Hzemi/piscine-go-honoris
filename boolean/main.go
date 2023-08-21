package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}
func main() {

	Even := "I have an even number of arguments"
	Odd := "I have an odd number of arguments"
	lengthOfArg := 0
	for range os.Args[1:] {
		lengthOfArg++
	}

	if isEven(lengthOfArg) {
		printStr(Even)
	} else {
		printStr(Odd)
	}
}
