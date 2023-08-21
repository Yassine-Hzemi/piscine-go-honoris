package piscine

func Compact(ptr *[]string) int {
	length := 0
	for _, c := range *ptr {
		if c != "" {
			length++
		}
	}
	count := 0
	arr_new := make([]string, length)
	for _, c := range *ptr {
		if c != "" {
			arr_new[count] = c
			count++
		}
	}
	*ptr = arr_new
	return length
}
