package piscine

// CompStr function
func CompStr(a, b interface{}) bool {
	return a == b
}

// ListFind function returns the address of the first node
// in the list l that is determined to be equal
// to ref by the function CompStr.
func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	found := &NodeL{}
	for l.Head != nil {
		if comp(l.Head.Data, ref) {
			return &l.Head.Data
		}
		l.Head = l.Head.Next
	}
	return &found.Data
}
