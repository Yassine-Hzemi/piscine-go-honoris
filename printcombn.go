package piscine

import (
	"github.com/01-edu/z01"
)

func printNumber(n int) {

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n >= 10 {
		printNumber(n / 10)
	}

	w := '0'

	for i := 0; i < n%10; i++ {
		w++
	}

	z01.PrintRune(w)
}

func isValid(n int) bool {
	for n != 0 {
		if n >= 10 && n%10 <= (n/10)%10 {
			return false
		}
		n /= 10
	}
	return true
}

func isLastNum(nb, pw int) bool {
	mxLn := 1
	for i := 1; i < pw; i++ {
		mxLn *= 10
	}

	if nb/mxLn == 10-pw {
		return true
	}

	return false
}

func PrintCombN(n int) {
	mxLn := 1
	for i := 1; i < n; i++ {
		mxLn *= 10
	}

	for i := mxLn / 10; i <= mxLn*9; i++ {
		if isValid(i) {
			if i <= mxLn && mxLn != 1 {
				z01.PrintRune('0')
			}

			printNumber(i)

			if !isLastNum(i, n) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
