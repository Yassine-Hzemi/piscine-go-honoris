package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	nd := root
	if nd == nil {
		return nil
	}
	for nd.Left != nil {
		nd = nd.Left
	}
	return nd
}
