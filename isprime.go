//package main
//import "fmt"

package piscine

//IsPrime shows if x is a prime number or no
func IsPrime(x int) bool {
	if x > 0 {
		if x == 1 {
			return false
		}
		for m := 2; m*m <= x; m++ {
			if x%m == 0 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

/*func main() {
fmt.Println(Fibonacci(7))
}*/
