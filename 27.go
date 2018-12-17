package algo

func mirrorRecursively(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		mirrorRecursively(root.Left)
	}
	if root.Right != nil {
		mirrorRecursively(root.Right)
	}
	return root
}
