package piscine

func MakeRange(min, max int) []int {
	if !(min >= max) {
		size := max - min
		numbersTable := make([]int, size)
		for i := 0; i < size; i++ {
			numbersTable[i] = min + i
		}
		return numbersTable
	} else {
		return []int(nil)
	}
}
