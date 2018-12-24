package algo

func deleteNode(list **LinkList, toBeDeleted *LinkList) {
	if *list == nil || toBeDeleted == nil {
		return
	}
	if toBeDeleted.Next != nil {
		next := toBeDeleted.Next
		toBeDeleted.Val = next.Val
		toBeDeleted.Next = next.Next
	} else if *list == toBeDeleted {
		*list = nil
		toBeDeleted = nil
	} else {
		node := *list
		for node.Next != toBeDeleted {
			node = node.Next
		}
		node.Next = nil
		toBeDeleted = nil
	}
}
