package algo

func isBalancedTree(root *TreeNode) (result bool) {
	result = false
	if root == nil {
		result = true
		return
	}
	left := treeDepth(root.Left)
	right := treeDepth(root.Right)
	diff := right - left
	if diff > 1 || diff < -1 {
		return
	}
	return isBalancedTree(root.Left) && isBalancedTree(root.Right)
}

func isBalancedTreeTwo(root *TreeNode) bool {
	depth := 0
	return isBalanced(root, &depth)
}

func isBalanced(root *TreeNode, depth *int) (result bool) {
	result = false
	if root == nil {
		*depth = 0
		result = true
		return
	}
	left, right := 0, 0
	if isBalanced(root.Left, &left) && isBalanced(root.Right, &right) {
		diff := right - left
		if diff >= -1 && diff <= 1 {
			if left > right {
				*depth = left - right + 1
			} else {
				*depth = right - left + 1
			}
			result = true
			return
		}
	}
	return
}
