//package main
package piscine

//import "fmt"
func StrRev(s string) string {
	Count := 0
	stringInBytes := []byte(s)
	for range s {
		Count++
	}

	Count--

	for index := range s {
		stringInBytes[index] = s[Count-index]
	}

	str := string(stringInBytes)

	return str
}

/*func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}*/
