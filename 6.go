package algo

func printLinkListReversinglyIteratively(list *LinkList) (result []int) {

	if list == nil {
		return
	}
	stack := make([]*LinkList, 0)
	node := list
	for node != nil {
		stack = append(stack, node)
		node = node.Next
	}
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		result = append(result, top.Val)
		stack = stack[:len(stack)-1]
	}
	return
}
