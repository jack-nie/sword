package algo

func findListKNums(nums []int, k int) (result []int) {
	result = nil
	if len(nums) == 0 {
		return
	}
	start := 0
	end := len(nums) - 1
	index := partition(nums, start, end)
	for index != k-1 {
		if index < k-1 {
			start = index + 1
			index = partition(nums, start, end)
		} else {
			end = index - 1
			index = partition(nums, start, end)
		}
	}
	for i := 0; i < k; i++ {
		result = append(result, nums[i])
	}
	return
}
