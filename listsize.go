package piscine

// ListSize function returns the number
// of elements in a linked list l.
func ListSize(l *List) int {
	num := 0
	for l.Head != nil {
		num++
		l.Head = l.Head.Next
	}
	return num
}
