package algo

func entryNodeOfLoop(list *LinkList) *LinkList {
	meeting := meetingNode(list)
	if meeting == nil {
		return nil
	}
	count := 1
	pHead := meeting
	for pHead.Next != meeting {
		pHead = pHead.Next
		count++
	}
	pNode1 := list
	pNode2 := list
	for i := 0; i < count; i++ {
		pNode1 = pNode1.Next
	}
	for pNode1 != pNode2 {
		pNode1 = pNode1.Next
		pNode2 = pNode2.Next
	}
	return pNode1
}

func meetingNode(list *LinkList) *LinkList {
	if list == nil {
		return nil
	}
	slow := list.Next
	if slow == nil {
		return nil
	}
	fast := slow.Next
	for fast != nil && slow != nil {
		if fast == slow {
			return fast
		}
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return nil
}
