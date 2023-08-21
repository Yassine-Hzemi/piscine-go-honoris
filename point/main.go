package main

import (
	"github.com/01-edu/z01"
)

type cord struct {
	x int
	y int
}

func setPoint(ptr *cord) {
	ptr.x = 42
	ptr.y = 21
}
func printS(s string) {
	for _, sym := range s {
		z01.PrintRune(sym)
	}
}

func intToRune(num int) {
	if num < 10 {
		for i, j := 0, '0'; i <= num%10; i, j = i+1, j+1 {
			if i == num%10 {
				z01.PrintRune(j)
			}
		}
	} else {
		nextnum := num / 10
		mod := num % 10
		intToRune(nextnum)
		intToRune(mod)
	}

}
func main() {
	pt := cord{}
	points := pt

	setPoint(&points)
	printS("x = ")
	//	printI(points.x)
	intToRune(points.x)

	printS(", y = ")
	//	printI(points.y)
	intToRune(points.y)

	z01.PrintRune('\n')
}
