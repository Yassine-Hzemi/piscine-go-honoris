//package main
//import "fmt"

package piscine

func Sqrt(x int) int {
	r := 0
	for i := 0; i <= x; i++ {
		if r*r == x {
			return r
		} else {
			r += 1
		}
	}
	return 0
}

/*func main() {
fmt.Println(Fibonacci(9))
}*/
