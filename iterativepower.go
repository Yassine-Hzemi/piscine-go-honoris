//package main
//import "fmt"
package piscine

// IterativePower factoriel
func IterativePower(x int, n int) int {
	if n < 0 {
		return 0
	}
	f := 1
	for n > 0 {
		f *= x
		n--
	}
	//fmt.Println(f)
	return f
}

/*func main() {
IterativePower(25,2)
}*/
