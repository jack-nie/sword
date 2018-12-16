package algo

func findKThNodeToTail(list *LinkList, k int) *LinkList {
	if list == nil {
		return nil
	}
	pHead := list
	var pBehind *LinkList
	for i := 0; i < k-1; i++ {
		if pHead.Next != nil {
			pHead = pHead.Next
		} else {
			return nil
		}
	}
	pBehind = list
	for pHead.Next != nil {
		pHead = pHead.Next
		pBehind = pBehind.Next
	}
	return pBehind
}
