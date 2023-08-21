//package main
//import "fmt"
package piscine

func LastRune(s string) rune {
	letter := []rune(s)
	last := 0
	for i := range s {
		last = i
	}
	return letter[last]
}

/*
func main() {
fmt.Printf("%c",LastRune("Hello!"))
fmt.Printf("%c" ,LastRune("Salut!"))
	fmt.Printf("%c",LastRune("Ola!"))
	fmt.Printf("\n")
}
*/
