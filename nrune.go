//package main
//import "fmt"
package piscine

func NRune(s string, n int) rune {
	letter := []rune(s)
	length := 0
	if n <= 0 {
		return 0
	}
	for i := range s {
		length = i + 1
	}
	if length == 0 || n > length {
		return 0
	}
	return letter[n-1]
}

/*
func main() {
fmt.Printf("%c",NRune("Hello!", 2))
fmt.Printf("%c" ,NRune("Salut!", 3))
fmt.Printf("%c",NRune("Ola!", 2))
fmt.Printf("\n")
}
*/
