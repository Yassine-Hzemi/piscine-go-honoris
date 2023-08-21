package piscine

func Unmatch(arr []int) int {
	for _, i := range arr {
		a := 0
		for _, l := range arr {
			if i == l {
				a++
			}
		}
		if a%2 == 1 {
			return i
		}
	}
	return -1
}
