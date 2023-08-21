package piscine

// ListPushFront function inserts a new element NodeL
// at the beginning of the list l while using the structure List.
func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	newNode.Next = l.Head
	l.Head = newNode
}
