package piscine

func ForEach(f func(int), arr []int) {
	for _, a := range arr {
		f(a)
	}
}
