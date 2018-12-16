package algo

func findMissingNums(nums []int, n int) int {
	if len(nums) == 0 || len(nums) != n {
		return -1
	}
	left := 0
	right := n - 1
	for left <= right {
		middle := (right-left)>>1 + left
		if nums[middle] != middle {
			if middle == 0 || nums[middle-1] == middle-1 {
				return middle
			}
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	if left == n {
		return n
	}
	return -1
}
