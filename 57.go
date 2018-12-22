package algo

func findNumsWithSum(nums []int, target int) (result []int) {
	if len(nums) == 0 {
		return nil
	}

	left := 0
	right := len(nums) - 1
	result = make([]int, 2)
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			result[0] = nums[left]
			result[1] = nums[right]
			return
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
