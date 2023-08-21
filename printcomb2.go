package piscine

import "github.com/01-edu/z01"

func PrintComb2() {

	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for h := '0'; h <= '9'; h++ {
				for k := '0'; k <= '9'; k++ {
					if i > h || (i == h && j >= k) {
						continue
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(32)
					z01.PrintRune(h)
					z01.PrintRune(k)
					if !(i == '9' && j == '8' && h == '9' && k == '9') {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}

				}
			}
		}
	}
	z01.PrintRune(10)
}
