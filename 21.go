package algo

func reorderOddEven(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	start := 0
	end := len(nums) - 1
	for start < end {
		for start < end && nums[start]&0x1 != 0 {
			start++
		}
		for start < end && nums[end]&0x1 == 0 {
			end--
		}
		if start < end {
			nums[start], nums[end] = nums[end], nums[start]
		}
	}
	return nums
}

func reorderOddEvenTwo(nums []int, f func(int) bool) []int {
	if len(nums) == 0 {
		return nil
	}
	start := 0
	end := len(nums) - 1
	for start < end {
		for start < end && isEven(nums[start]) {
			start++
		}
		for start < end && !isEven(nums[end]) {
			end--
		}
		if start < end {
			nums[start], nums[end] = nums[end], nums[start]
		}
	}
	return nums
}

func isEven(i int) bool {
	if i&0x1 != 0 {
		return true
	}
	return false
}
