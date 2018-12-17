package algo

func hasSubTree(root1, root2 *TreeNode) bool {
	result := false

	if root1 != nil && root2 != nil {
		if root1.Val == root2.Val {
			result = doseTree1HasTree2(root1, root2)
		}
		if !result {
			result = hasSubTree(root1.Left, root2)
		}
		if !result {
			result = hasSubTree(root1.Right, root2)
		}
	}
	return result
}

func doseTree1HasTree2(root1, root2 *TreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return doseTree1HasTree2(root1.Left, root2.Left) && doseTree1HasTree2(root1.Right, root2.Right)
}
