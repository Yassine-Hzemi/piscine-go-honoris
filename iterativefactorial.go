//package main
//import "fmt"

package piscine

// IterativeFactorial factoriel
func IterativeFactorial(x int) int {
	f := 1
	if x > 20 {
		f = 0
	} else if x > 0 {
		for x > 0 {
			f *= x
			x--
		}
	} else if x == 0 {
		f = 1
	} else {
		f = 0
	}
	//fmt.Println(f)
	return f
}

/*
func main() {
IterativeFactorial(0)
}*/
