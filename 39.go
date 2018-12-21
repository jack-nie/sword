package algo

func moreThanHalfNums(nums []int) (result int) {
	if len(nums) == 0 {
		return -1
	}
	result = nums[0]
	times := 1
	for i := 1; i < len(nums); i++ {
		if times == 0 {
			result = nums[i]
			times = 1
		} else if nums[i] == result {
			times++
		} else {
			times--
		}
	}
	if !checkMoreThanHalf(nums, result) {
		return -1
	}
	return
}

func checkMoreThanHalf(nums []int, target int) (result bool) {
	times := 0
	result = false
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			times++
		}
	}
	if times*2 > len(nums) {
		return true
	}
	return
}
