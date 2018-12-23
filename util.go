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
