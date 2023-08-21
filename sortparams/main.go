package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	length := 0
	for range args {
		length++
	}

	for i := 1; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}

	for i := 1; i < length; i++ {
		for _, letter := range args[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
