package piscine

func ActiveBits(n int) (active int) {
	for n != 0 {
		if n%2 == 1 {
			active++
		}
		n /= 2
	}
	return
}
