package algo

func findFirstCommonNode(list1, list2 *LinkList) *LinkList {
	list1Length := getLengthOfList(list1)
	list2Length := getLengthOfList(list2)
	diff := list1Length - list2Length
	longHead := list1
	shortHead := list2
	if list2Length > list1Length {
		longHead = list2
		shortHead = list1
		diff = list2Length - list1Length
	}
	for i := 0; i < diff; i++ {
		longHead = longHead.Next
	}
	for longHead != nil && shortHead != nil && longHead != shortHead {
		longHead = longHead.Next
		shortHead = shortHead.Next
	}
	return longHead
}

func getLengthOfList(list *LinkList) int {
	if list == nil {
		return 0
	}
	length := 0
	for list != nil {
		list = list.Next
		length++
	}
	return length
}
