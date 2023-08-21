package piscine

func AppendRange(min, max int) []int {
	var numbersTable []int
	if !(min >= max) {
		for i := min; i < max; i++ {
			numbersTable = append(numbersTable, i)
		}
	}
	return numbersTable
}
