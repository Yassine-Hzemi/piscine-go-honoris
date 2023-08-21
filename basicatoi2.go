//package main

//import "fmt"

package piscine

//import "fmt"

func BasicAtoi2(s string) int {
	byteRevStr := []byte(s)
	lengthStr := -1

	for range s {
		lengthStr++
	}

	for index := range s {
		byteRevStr[lengthStr-index] = s[index]
		byteRevStr[lengthStr-index] -= 48
	}

	num := 0
	tenth := 1
	//fmt.Println(byteRevStr)
	for index := range s {
		if int(byteRevStr[index]) > 9 || int(byteRevStr[index]) < 0 {
			return 0
		}
		num = num + (int(byteRevStr[index]) * tenth)
		tenth *= 10
	}
	return num
}

/*
func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"

	n := BasicAtoi2(s)
	n2 := BasicAtoi2(s2)
	n3 := BasicAtoi2(s3)
	n4 := BasicAtoi2(s4)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

}
*/
