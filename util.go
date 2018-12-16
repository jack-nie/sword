package algo

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

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
