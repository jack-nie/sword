package algo

func maxProfit(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	min := nums[0]
	maxDiff := nums[1] - min

	for i := 2; i < len(nums); i++ {
		if nums[i-1] < min {
			min = nums[i-1]
		}
		currentDiff := nums[i] - min
		if currentDiff > maxDiff {
			maxDiff = currentDiff
		}
	}
	return maxDiff
}
