package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var n rune
	var m rune
	var o rune
	for n = 0; n <= 9; n++ {
		for m = 0; m <= 9; m++ {
			if m != n && m > n {
				for o = 0; o <= 9; o++ {
					if o != m && o > m {
						z01.PrintRune(n + 48)
						z01.PrintRune(m + 48)
						z01.PrintRune(o + 48)
						if !(n == 7 && m == 8 && o == 9) {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
