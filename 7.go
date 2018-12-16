package algo

func constructBinaryTree(preOrder, inOrder []int) *TreeNode {
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}

	preOrderLength := len(preOrder)
	inOrderLength := len(inOrder)
	return constructCore(preOrder, inOrder, &preOrderLength, &inOrderLength)

}

func constructCore(preOrder, inOrder []int, preOrderLength, inOrderLength *int) *TreeNode {
	root := new(TreeNode)
	root.Val = preOrder[0]
	return root
}
