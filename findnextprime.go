//package main
//import "fmt"

package piscine

//FindNextPrime 1st prime after x and x
func FindNextPrime(x int) int {
	if x > 0 {
		if x <= 1 {
			return 2
		}
		for m := 2; m*m <= x; m++ {
			if x%m == 0 {
				return FindNextPrime(x + 1)
			}
		}
		return x
	} else {
		return 2
	}
}

/*
func main() {
fmt.Println(FindNextPrime(10))
}*/
