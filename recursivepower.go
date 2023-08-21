//package main
//import "fmt"
package piscine

// RecursivePower factoriel
func RecursivePower(x int, n int) int {
	f := 1
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n < 0 {
		return 0
	}
	f = x * RecursivePower(x, n-1)
	//	fmt.Println(f)
	return f
}

/*
func main() {
RecursivePower(25,2)
}*/
