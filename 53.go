package algo

func findNumsOfK(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	first := getFirstOfK(nums, k, 0, len(nums)-1)
	last := getLastOfK(nums, k, 0, len(nums)-1)
	if first > -1 && last > -1 {
		return last - first + 1
	}
	return 0
}

func getFirstOfK(nums []int, k, left, right int) int {
	if left > right {
		return -1
	}
	middleIndex := (right-left)>>1 + left
	middleData := nums[middleIndex]
	if middleData == k {
		if middleIndex > 0 && nums[middleIndex-1] != k || middleIndex == 0 {
			return middleIndex
		}
		right = middleIndex - 1
	} else if middleData > k {
		right = middleIndex - 1
	} else {
		left = middleIndex + 1
	}
	return getFirstOfK(nums, k, left, right)
}

func getLastOfK(nums []int, k, left, right int) int {
	if left > right {
		return -1
	}
	middleIndex := (right-left)>>1 + left
	middleData := nums[middleIndex]
	if middleData == k {
		if middleIndex < len(nums)-1 && nums[middleIndex+1] != k || middleIndex == len(nums)-1 {
			return middleIndex
		}
		left = middleIndex + 1
	} else if middleData > k {
		right = middleIndex - 1
	} else {
		left = middleIndex + 1
	}
	return getLastOfK(nums, k, left, right)
}
