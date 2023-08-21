//package main
//import "fmt"
package piscine

func FirstRune(s string) rune {
	letter := []rune(s)
	length := 0
	for i := range s {
		length = i + 1
	}
	if length == 0 {
		return -1
	}
	return letter[0]

}

/*func main() {
fmt.Printf("%c",FirstRune("Hello!"))
fmt.Printf("%c" ,FirstRune("Salut!"))
	fmt.Printf("%c",FirstRune("Ola!"))
	fmt.Printf("\n")
}
*/
