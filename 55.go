package algo

func treeDepth(root *TreeNode) (result int) {
	result = 0
	if root == nil {
		return
	}
	left := treeDepth(root.Left)
	right := treeDepth(root.Right)
	if left >= right {
		return left + 1
	}
	return right + 1
}
