package algo

func findNumSameAsIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := (right-left)>>1 + left
		if nums[middle] == middle {
			return middle
		} else if nums[middle] > middle {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}
