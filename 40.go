package algo

func findListKNums(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
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
	var out []int
	for i := 0; i < k; i++ {
		out = append(out, nums[i])
	}
	return out
}
