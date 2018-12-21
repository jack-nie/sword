package algo

func findInRotateArray(nums []int) (result int) {
	if len(nums) == 0 {
		return -1
	}
	index1 := 0
	index2 := len(nums) - 1
	indexMid := index1
	for nums[index1] >= nums[index2] {
		if index2-index1 == 1 {
			indexMid = index2
			break
		}

		indexMid = (index1 + index2) / 2
		if nums[index1] == nums[index2] && nums[indexMid] == nums[index1] {
			return minInOrder(nums, index1, index2)
		}

		if nums[indexMid] >= nums[index1] {
			index1 = indexMid
		} else if nums[indexMid] <= nums[index2] {
			index2 = indexMid
		}
	}
	result = nums[indexMid]
	return
}

func minInOrder(nums []int, index1, index2 int) (result int) {
	result = nums[index1]
	for i := index1 + 1; i <= index2; i++ {
		if result > nums[i] {
			result = nums[i]
		}
	}
	return
}
