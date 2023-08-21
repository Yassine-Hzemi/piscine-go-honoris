package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	nd := root
	if nd == nil {
		return nil
	}
	for nd.Right != nil {
		nd = nd.Right
	}
	return nd
}
