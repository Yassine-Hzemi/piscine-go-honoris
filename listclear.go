package piscine

// ListClear function deletes all nodes from a linked list l.
func ListClear(l *List) {
	if l.Head != nil || l.Tail != nil {
		l.Head = nil
		l.Tail = nil
	}
}
