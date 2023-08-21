//package main
//import "fmt"
package piscine

// RecursiveFactorial factoriel
func RecursiveFactorial(x int) int {
	f := 1
	if x > 20 {
		f = 0
	} else if x > 0 {
		f = x * RecursiveFactorial(x-1)
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
RecursiveFactorial(0)
}
*/
