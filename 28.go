package algo

func isSymmetrical(root *TreeNode) bool {
	return checkIsSymmetrical(root, root)
}

func checkIsSymmetrical(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return checkIsSymmetrical(root1.Left, root2.Right) &&
		checkIsSymmetrical(root1.Right, root2.Left)
}
