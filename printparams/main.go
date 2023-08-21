package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	i := 2
	for i <= len(os.Args) {
		z01.PrintRune("%+v", os.Args[i-1])
		z01.PrintRune(10)
		i++
	}
}
