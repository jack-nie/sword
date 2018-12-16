package algo

func mergeTwoSortedList(list1, list2 *LinkList) *LinkList {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var pMergedHead *LinkList
	if list1.Val > list2.Val {
		pMergedHead = list2
		pMergedHead.Next = mergeTwoSortedList(list1, list2.Next)
	} else {
		pMergedHead = list1
		pMergedHead.Next = mergeTwoSortedList(list1.Next, list2)
	}
	return pMergedHead
}
