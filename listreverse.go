package piscine

// ListReverse function reverses the order
// of the elements of a given linked list l.
func ListReverse(l *List) {
	lnew := &List{}
	for l.Head != nil {
		ListPushFront(lnew, l.Head.Data)
		l.Head = l.Head.Next
	}
	l.Head, l.Tail = lnew.Head, lnew.Tail
}
