package piscine

func Max(arr []int) int {
	max := 0
	for _, n := range arr {
		if n > max {
			max = n
		}
	}
	return max
}
