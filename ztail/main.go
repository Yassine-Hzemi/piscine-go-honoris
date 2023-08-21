package main

import (
	"fmt"
	"os"
)

const (
	success = iota
	failure
)

var status = success

func notNil(err error) bool {
	if err != nil {
		status = failure
		fmt.Printf("%s\n", err)
		return true
	}
	return false
}
func length(a []string) (i int) {
	for range a {
		i++
	}
	return
}
func main() {
	bytes := atoi(os.Args[2])

	filenames := os.Args[3:]
	for i, filename := range filenames {
		file, err := os.Open(filename)
		if notNil(err) {
			continue
		}
		defer file.Close()
		fileInfo, err := file.Stat()
		if notNil(err) {
			continue
		}
		offset := fileInfo.Size() - int64(bytes)
		if offset < 0 {
			offset = 0
		}
		var b []byte
		for i := int64(0); i < fileInfo.Size()-offset; i++ {
			b = append(b, 0)
		}
		_, err = file.ReadAt(b, offset)
		if notNil(err) {
			continue
		}
		if length(filenames) > 1 {
			if i > 0 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> " + filename + " <==\n")
		}
		os.Stdout.Write(b)
	}
	os.Exit(status)
}

func atoi(s string) int {
	str := 0
	signCount := 0
	addMinus := false
	for i, v := range s {
		if int(rune(v)) > 47 && int(rune(v)) < 58 || v == '+' || v == '-' {
			if (v == '-' || v == '+') && i > 0 {
				return 0
			}
			if v == '-' && i == 0 {
				addMinus = true
			}
			if v == '+' || v == '-' {
				signCount++
				if signCount > 1 {
					return 0
				}
				continue
			}
			a := 0
			for i := '1'; i <= v; i++ {
				a++
			}
			str = str*10 + a
		} else {
			return 0
		}
	}
	if addMinus {
		str = str - (str * 2)
	}
	return str
}
