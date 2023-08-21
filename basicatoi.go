//package main
package piscine

//import "fmt"

func BasicAtoi(s string) int {
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
		num = num + (int(byteRevStr[index]) * tenth)
		tenth *= 10
	}
	return num
}

/*
func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "000000"

	n := BasicAtoi(s)
	n2 := BasicAtoi(s2)
	n3 := BasicAtoi(s3)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
}*/
