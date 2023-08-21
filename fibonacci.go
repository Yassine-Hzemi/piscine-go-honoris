//package main
//import "fmt"

package piscine

func Fibonacci(x int) int {
	if x < 0 {
		return -1
	}
	if x <= 1 {
		return x
	}
	return Fibonacci(x-1) + Fibonacci(x-2)
}

/*func main() {
fmt.Println(Fibonacci(9))
}*/
