package piscine

// ListLast function returns the last element of a linked list l.
func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current.Data
}
