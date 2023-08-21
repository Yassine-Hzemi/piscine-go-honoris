package piscine

// SortWordArr function sorts a string array by ascii  in asc order.
func SortWordArr(array []string) {
	lenarray := 0
	for i := range array {
		lenarray = i + 1
	}
	for i := 0; i < lenarray-1; i++ {
		for j := i + 1; j < lenarray; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
