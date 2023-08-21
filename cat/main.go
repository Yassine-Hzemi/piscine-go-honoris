package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func put(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
func putln(s string) {
	put(s)
	z01.PrintRune('\n')
}

func fatal(s string) {
	putln(s)
	os.Exit(1)
}

func main() {
	i := 0
	for range os.Args {
		i++
	}
	if i == 1 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fatal(err.Error())
		}
		put(string(b))
	} else {
		for _, file := range os.Args[1:] {
			b, err := ioutil.ReadFile(file)
			if err != nil {
				fatal("ERROR: " + err.Error())
			}
			put(string(b))
			z01.PrintRune('\n')

		}
	}
}
