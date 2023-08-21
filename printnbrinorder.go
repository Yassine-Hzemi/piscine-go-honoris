package piscine

import (
	"github.com/01-edu/z01"
)

func SortTable(table []int) {
	l := 0
	i := 1
	for range table {
		l++
	}
	for i < l {
		if table[i-1] > table[i] {
			tmp := table[i]
			table[i] = table[i-1]
			table[i-1] = tmp
			i = 1
		} else {
			i++
		}
	}
}

func PrintNbrInOrder(n int) {
	if n > 0 {
		var s []int
		for n != 0 {
			a := n % 10
			n /= 10
			s = append(s, a)
		}
		SortTable(s)
		for _, a := range s {
			z01.PrintRune(rune(a + 48))
		}
	} else if n == 0 {
		z01.PrintRune('0')
	}
}
