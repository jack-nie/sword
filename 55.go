package algo

func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := treeDepth(root.Left)
	right := treeDepth(root.Right)
	if left >= right {
		return left + 1
	} else {
		return right + 1
	}
}
