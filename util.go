package algo

// TreeNode represents the tree structure
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// LinkList represent the link list structure
type LinkList struct {
	Next *LinkList
	Val  int
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func doubleEqual(a, b float64) bool {
	if a-b > -0.00000001 && a-b < 0.00000001 {
		return true
	}
	return false
}

func generateListFromSlice(nums []int) *LinkList {
	list := new(LinkList)
	node := list
	for i := 0; i < len(nums); i++ {
		listPushBack(list, nums[i])
	}
	return node.Next
}

func listPushBack(list *LinkList, value int) {
	newNode := new(LinkList)
	newNode.Val = value
	newNode.Next = list.Next
	list.Next = newNode
}

func convertListToSlice(list *LinkList) []int {
	if list == nil {
		return nil
	}

	node := list
	result := make([]int, 0)
	for node.Next != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}
