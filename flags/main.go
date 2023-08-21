package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func order(s string) {
	var a [5000]int
	for _, c := range s {
		a[int(c)]++
	}
	for i, c := range a {
		for c > 0 {
			z01.PrintRune(rune(i))
			c--
		}
	}
	z01.PrintRune('\n')
}

func insert(str1 string, str2 string) string {
	return str1 + str2
}

func help() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func main() {

	myArr := os.Args[1:]
	ln := -1

	for i := range myArr {
		ln = i
	}
	if ln != -1 {
		if myArr[0] == "-h" || myArr[0] == "--help" {
			help()

		} else if myArr[0] == "--order" || myArr[0] == "-o" {
			order(myArr[1])

		} else if myArr[0][0:3] == "--i" || myArr[0][0:2] == "-i" {
			if ln < 2 {
				if myArr[0][0:3] == "--i" {
					fmt.Println(insert(myArr[1], myArr[0][9:]))
				} else {
					fmt.Println(insert(myArr[1], myArr[0][3:]))
				}
			} else {
				myStr := ""
				for i := 1; i <= ln; i++ {
					if myArr[i] != "-o" && myArr[i] != "--order" {
						myStr = insert(myStr, myArr[i])
					}
				}

				if myArr[0][0:3] == "--i" {
					myStr = insert(myStr, myArr[0][9:])
				} else {
					myStr = insert(myStr, myArr[0][3:])
				}
				order(myStr)
			}
		} else {
			fmt.Println(myArr[0])
		}
	} else {
		help()
	}
}
