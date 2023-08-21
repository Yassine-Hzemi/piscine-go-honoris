package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	n := l.Head
	prev := n

	for n != nil && n.Data == data_ref {
		l.Head = n.Next
		n = l.Head
	}
	for n != nil {
		for n != nil && n.Data != data_ref {
			prev = n
			n = n.Next
		}
		if n == nil {
			return
		}
		prev.Next = n.Next
		n = n.Next
	}
}
