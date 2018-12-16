package algo

func reverseList(list *LinkList) *LinkList {
	if list == nil {
		return nil
	}
	var pReversedHead, prev *LinkList
	pNode := list
	for pNode != nil {
		pNext := pNode.Next
		if pNext == nil {
			pReversedHead = pNode
		}
		pNode.Next = prev
		prev = pNode
		pNode = pNext
	}
	return pReversedHead
}
