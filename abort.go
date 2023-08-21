package piscine

//hahahahahahaha
func Abort(a, b, c, d, e int) int {
	median := SortArray(a, b, c, d, e)[2]
	return median
}

func SortArray(a, b, c, d, e int) []int {
	arr := []int{a, b, c, d, e}
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
