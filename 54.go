package algo

func findKthNode(root *TreeNode, k *int) (target *TreeNode) {
	if root == nil || *k < 0 {
		return nil
	}
	if root.Left != nil {
		target = findKthNode(root.Left, k)
	}

	if target == nil {
		if *k == 1 {
			target = root
		}
		*k--
	}

	if target == nil && root.Right != nil {
		target = findKthNode(root.Right, k)
	}
	return
}
