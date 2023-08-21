package piscine

func myStrLen(s string) int {
	length := 0
	for range s {
		length++
	}
	return length
}

func Split(str, charset string) []string {
	strLen := myStrLen(str)
	charsetLen := myStrLen(charset)
	arrLen := 0
	index := 0
	for i1, v1 := range str {
		checker := 0
		if byte(v1) == charset[0] && strLen-i1 >= charsetLen {
			for i2, v2 := range charset {
				if byte(v2) == str[i1+i2] {
					checker++
				}
			}
			if checker == charsetLen {
				arrLen++
			}
		}
	}

	strArr := make([]string, arrLen+1)

	for i1 := 0; i1 < strLen; i1++ {
		checker := 0
		if str[i1] == charset[0] && strLen-i1 >= charsetLen {
			for i2, v2 := range charset {
				if byte(v2) == str[i1+i2] {
					checker++
				}
			}
			if checker == charsetLen {
				i1 += charsetLen - 1
				index++
			} else {
				strArr[index] += string(str[i1])
			}
		} else {
			strArr[index] += string(str[i1])
		}
	}
	return strArr
}
