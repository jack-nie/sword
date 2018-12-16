package algo

func findDup(nums []int, n int) int {
	if len(nums) == 0 {
		return -1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] > n-1 {
			return -1
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
