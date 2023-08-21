package piscine

func SortIntegerTable(table []int) {
	var temp int
	var length int = 0

	for range table {
		length++
	}

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if table[j] > table[j+1] {
				temp = table[j+1]
				table[j+1] = table[j]
				table[j] = temp
			}
		}
	}
}
