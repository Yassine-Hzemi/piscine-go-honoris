package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	length := 0
	for i := range tab {
		length = i + 1
	}
	left := false
	right := false
	for i := 0; i < length-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			left = true
		} else if f(tab[i], tab[i+1]) < 0 {
			right = true
		}
	}
	if left && right {
		return false
	} else {
		return true
	}
}
