package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func convertBinaryTreeToDoubleLinkList(root *TreeNode) *TreeNode {
	var pLastNodeInList *TreeNode
	convertNode(root, &pLastNodeInList)
	pHeadOfList := pLastNodeInList

	for pHeadOfList != nil && pHeadOfList.Left != nil {
		pHeadOfList = pHeadOfList.Left
	}
	return pHeadOfList
}

func convertNode(root *TreeNode, pLastNodeInList **TreeNode) {
	if root == nil {
		return
	}
	pCurrent := root
	if pCurrent.Left != nil {
		convertNode(root.Left, pLastNodeInList)
	}
	pCurrent.Left = *pLastNodeInList
	if *pLastNodeInList != nil {
		(*pLastNodeInList).Right = pCurrent
	}
	*pLastNodeInList = pCurrent
	if pCurrent.Right != nil {
		convertNode(pCurrent.Right, pLastNodeInList)
	}
}

func main() {

}
