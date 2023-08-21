package piscine

//thank you for for all the bugs and waist of time
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		var tmp *TreeNode
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left == nil {
			tmp = root
			root = root.Right
			tmp = nil
		} else if root.Right == nil {
			tmp = root
			root = root.Left
			tmp = nil
		} else {
			tmp = BTreeMin(root.Right)
			root.Data = tmp.Data
			root.Right = BTreeDeleteNode(root.Right, node)
		}
	}
	return root
}
