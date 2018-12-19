package algo

func constructBinaryTree(preOrder, inOrder []int) *TreeNode {
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}

	preOrderLength := len(preOrder)
	inOrderLength := len(inOrder)
	return constructCore(preOrder, inOrder, 0, preOrderLength-1, 0, inOrderLength-1)

}

func constructCore(preOrder, inOrder []int, ps, pe, is, ie int) *TreeNode {
	if ps > pe {
		return nil
	}
	value := preOrder[ps]
	index := is
	for index <= ie && inOrder[index] != value {
		index++
	}
	root := new(TreeNode)
	root.Val = value

	root.Left = constructCore(preOrder, inOrder, ps+1, ps+index-is, is, index-1)
	root.Right = constructCore(preOrder, inOrder, ps+index-is+1, pe, index+1, ie)
	return root
}
